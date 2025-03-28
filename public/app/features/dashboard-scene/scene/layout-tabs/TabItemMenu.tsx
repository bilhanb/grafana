import { css } from '@emotion/css';

import { GrafanaTheme2 } from '@grafana/data';
import { Button, Dropdown, Menu, ToolbarButtonRow, useStyles2 } from '@grafana/ui';
import { t, Trans } from 'app/core/internationalization';

import { TabItem } from './TabItem';

interface Props {
  model: TabItem;
}

export function TabItemMenu({ model }: Props) {
  const styles = useStyles2(getStyles);

  return (
    <ToolbarButtonRow className={styles.container}>
      <Dropdown
        placement="bottom-end"
        overlay={() => (
          <Menu>
            <Menu.Item
              ariaLabel={t('dashboard.tabs-layout.tab.menu.add-panel', 'Panel')}
              label={t('dashboard.tabs-layout.tab.menu.add-panel', 'Panel')}
              onClick={() => model.onAddPanel()}
            />
            <Menu.Divider />
            <Menu.Item
              ariaLabel={t('dashboard.tabs-layout.tab.menu.add-tab-before', 'Tab before')}
              label={t('dashboard.tabs-layout.tab.menu.add-tab-above', 'Tab before')}
              onClick={() => model.onAddTabBefore()}
            />
            <Menu.Item
              ariaLabel={t('dashboard.tabs-layout.tab.menu.add-tab-after', 'Tab after')}
              label={t('dashboard.tabs-layout.tab.menu.add-tab-after', 'Tab after')}
              onClick={() => model.onAddTabAfter()}
            />
          </Menu>
        )}
      >
        <Button
          aria-label={t('dashboard.tabs-layout.tab.menu.add', 'Add tab')}
          title={t('dashboard.tabs-layout.tab.menu.add', 'Add tab')}
          tooltip={t('dashboard.tabs-layout.tab.menu.add', 'Add tab')}
          icon="plus"
          variant="secondary"
          size="sm"
        >
          <Trans i18nKey="grafana-ui.tags-input.add">Add</Trans>
        </Button>
      </Dropdown>
      <Dropdown
        placement="bottom-end"
        overlay={() => (
          <Menu>
            <Menu.Item
              aria-label={t('dashboard.tabs-layout.tab.menu.move-left', 'Move tab left')}
              label={t('dashboard.tabs-layout.tab.menu.move-left', 'Move tab left')}
              onClick={() => model.onMoveLeft()}
            />
            <Menu.Divider />
            <Menu.Item
              aria-label={t('dashboard.tabs-layout.tab.menu.move-right', 'Move tab right')}
              label={t('dashboard.tabs-layout.tab.menu.move-right', 'Move tab right')}
              onClick={() => model.onMoveRight()}
            />
          </Menu>
        )}
      >
        <Button
          aria-label={t('dashboard.tabs-layout.menu.move-tab', 'Move tab')}
          title={t('dashboard.tabs-layout.menu.move-tab', 'Move tab')}
          tooltip={t('dashboard.tabs-layout.menu.move-tab', 'Move tab')}
          icon="arrows-h"
          variant="secondary"
          size="sm"
        />
      </Dropdown>
    </ToolbarButtonRow>
  );
}

function getStyles(theme: GrafanaTheme2) {
  return {
    container: css({
      gap: theme.spacing(1),
      flexShrink: 0,
    }),
  };
}
