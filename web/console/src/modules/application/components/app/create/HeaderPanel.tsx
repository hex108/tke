import * as React from 'react';
import { RootProps } from '../AppContainer';
import { Justify, Icon } from '@tencent/tea-component';
import { t, Trans } from '@tencent/tea-app/lib/i18n';

export class HeaderPanel extends React.Component<RootProps, {}> {
  goBack = () => {
    history.back();
  };

  render() {
    let title = t('新建应用');

    return (
      <Justify
        left={
          <React.Fragment>
            <div className="secondary-title">
              <a href="javascript:;" className="back-link" onClick={this.goBack.bind(this)}>
                <Icon type="btnback" />
                {t('返回')}
              </a>
              <span className="line-icon">|</span>
              <h2>{title}</h2>
            </div>
          </React.Fragment>
        }
      />
    );
  }
}
