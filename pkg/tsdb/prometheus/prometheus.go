package prometheus

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	sdkhttpclient "github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana/pkg/infra/httpclient"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/plugins"
	"github.com/grafana/grafana/pkg/plugins/backendplugin/coreplugin"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/grafana/grafana/pkg/tsdb/intervalv2"
	"github.com/prometheus/client_golang/api"
	apiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

var (
	plog         = log.New("tsdb.prometheus")
	legendFormat = regexp.MustCompile(`\{\{\s*(.+?)\s*\}\}`)
	safeRes      = 11000
)

const pluginID = "prometheus"

type Service struct {
	intervalCalculator intervalv2.Calculator
	im                 instancemgmt.InstanceManager
}

type DatasourceJSON struct {
	TimeInterval          string `json:"timeInterval"`
	HttpMethod            string `json:"httpMethod"`
	CustomQueryParameters string `json:"customQueryParameters"`
}

func ProvideService(cfg *setting.Cfg, httpClientProvider httpclient.Provider, pluginStore plugins.Store) (*Service, error) {
	plog.Debug("initializing")
	im := datasource.NewInstanceManager(newInstanceSettings(httpClientProvider))

	s := &Service{
		intervalCalculator: intervalv2.NewCalculator(),
		im:                 im,
	}

	factory := coreplugin.New(backend.ServeOpts{
		QueryDataHandler: s,
	})
	resolver := plugins.CoreDataSourcePathResolver(cfg, pluginID)
	if err := pluginStore.AddWithFactory(context.Background(), pluginID, factory, resolver); err != nil {
		plog.Error("Failed to register plugin", "error", err)
		return nil, err
	}

	return s, nil
}

func newInstanceSettings(httpClientProvider httpclient.Provider) datasource.InstanceFactoryFunc {
	return func(settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
		var jsonData DatasourceJSON
		err := json.Unmarshal(settings.JSONData, &jsonData)
		if err != nil {
			return nil, fmt.Errorf("error reading settings: %w", err)
		}
		httpCliOpts, err := settings.HTTPClientOptions()
		if err != nil {
			return nil, fmt.Errorf("error getting http options: %w", err)
		}

		// Set SigV4 service namespace
		if httpCliOpts.SigV4 != nil {
			httpCliOpts.SigV4.Service = "aps"
		}

		forceHttpGet := strings.ToLower(jsonData.HttpMethod) == "get"

		client, err := createClient(settings.URL, httpCliOpts, httpClientProvider, jsonData.CustomQueryParameters, forceHttpGet)
		if err != nil {
			return nil, err
		}

		mdl := DatasourceInfo{
			ID:           settings.ID,
			URL:          settings.URL,
			TimeInterval: jsonData.TimeInterval,
			promClient:   client,
		}

		return mdl, nil
	}
}

func (s *Service) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	if len(req.Queries) == 0 {
		return &backend.QueryDataResponse{}, fmt.Errorf("query contains no queries")
	}

	q := req.Queries[0]
	dsInfo, err := s.getDSInfo(req.PluginContext)
	if err != nil {
		return nil, err
	}

	var result *backend.QueryDataResponse
	switch q.QueryType {
	case "timeSeriesQuery":
		fallthrough
	default:
		result, err = s.executeTimeSeriesQuery(ctx, req, dsInfo)
	}

	return result, err
}

func createClient(url string, httpOpts sdkhttpclient.Options, clientProvider httpclient.Provider, customQueryParams string, forceHttpGet bool) (apiv1.API, error) {
	var middlewares []sdkhttpclient.Middleware
	if customQueryParams != "" {
		middlewares = append(middlewares, customQueryParametersMiddleware(plog, customQueryParams))

	}
	if forceHttpGet {
		middlewares = append(middlewares, forceHttpGetMiddleware(plog))
	}
	httpOpts.Middlewares = middlewares

	roundTripper, err := clientProvider.GetTransport(httpOpts)
	if err != nil {
		return nil, err
	}

	cfg := api.Config{
		Address:      url,
		RoundTripper: roundTripper,
	}

	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return apiv1.NewAPI(client), nil
}

func (s *Service) getDSInfo(pluginCtx backend.PluginContext) (*DatasourceInfo, error) {
	i, err := s.im.Get(pluginCtx)
	if err != nil {
		return nil, err
	}

	instance := i.(DatasourceInfo)

	return &instance, nil
}

// IsAPIError returns whether err is or wraps a Prometheus error.
func IsAPIError(err error) bool {
	// Check if the right error type is in err's chain.
	var e *apiv1.Error
	return errors.As(err, &e)
}

func ConvertAPIError(err error) error {
	var e *apiv1.Error
	if errors.As(err, &e) {
		return fmt.Errorf("%s: %s", e.Msg, e.Detail)
	}
	return err
}
