import { css } from '@emotion/react';
import { media, mediaQuery } from '../../../styles';
import HeaderLogo from './HeaderLogo';
import HeaderWebNavigation from './HeaderWebNavigation';
import HeaderUserNotification from '../HeaderUserNotification';
import HeaderUserMenu from '../HeaderUserMenu';
import ButtonThemeChange from '../../common/ButtonThemeChange';

interface HeaderWebProps {}

function HeaderWeb(props: HeaderWebProps) {
  return (
    <div css={block}>
      <HeaderLogo />
      <div css={itemBlock}>
        <HeaderWebNavigation />
        <div css={iconItemsBlock}>
          <HeaderUserNotification />
          <HeaderUserMenu />
          <ButtonThemeChange />
        </div>
      </div>
    </div>
  );
}

const block = css`
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 3.25rem;
  ${media.xxxsmall} {
    display: none;
  }
  ${media.small} {
    width: 93%;
    display: flex;
  }
  ${media.medium} {
    width: 96%;
  }
  ${mediaQuery(1285)} {
    max-width: 1250px;
  }
`;

const itemBlock = css`
  display: flex;
  justify-content: center;
`;

const iconItemsBlock = css`
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1.5rem;
  padding: 0 1rem;
`;

export default HeaderWeb;
