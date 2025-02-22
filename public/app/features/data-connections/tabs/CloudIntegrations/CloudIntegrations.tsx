import { css } from '@emotion/css';
import React, { ReactElement } from 'react';

import { GrafanaTheme2 } from '@grafana/data';
import { useStyles2 } from '@grafana/ui';
import { AppPluginLoader } from 'app/features/plugins/components/AppPluginLoader';

import { CLOUD_ONBOARDING_APP_ID } from '../../constants';

export function CloudIntegrations(): ReactElement | null {
  const s = useStyles2(getStyles);

  return (
    <div className={s.container}>
      <AppPluginLoader id={CLOUD_ONBOARDING_APP_ID} />
    </div>
  );
}

const getStyles = (theme: GrafanaTheme2) => ({
  // We would like to force the app to stay inside the provided tab
  container: css`
    position: relative;
  `,
});
