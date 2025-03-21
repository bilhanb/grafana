import { css } from '@emotion/css';
import { ReactNode } from 'react';
import Skeleton from 'react-loading-skeleton';

import { GrafanaTheme2 } from '@grafana/data';
import { useStyles2 } from '@grafana/ui';

export interface ScopesTreeLoadingProps {
  children: ReactNode;
  nodeLoading: boolean;
}

export function ScopesTreeLoading({ children, nodeLoading }: ScopesTreeLoadingProps) {
  const styles = useStyles2(getStyles);

  if (nodeLoading) {
    return <Skeleton count={5} className={styles.loader} />;
  }

  return children;
}

const getStyles = (theme: GrafanaTheme2) => {
  return {
    loader: css({
      margin: theme.spacing(0.5, 0),
    }),
  };
};
