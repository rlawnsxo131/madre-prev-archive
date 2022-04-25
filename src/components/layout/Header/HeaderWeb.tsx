import { css } from '@emotion/react';
import { media, mediaQuery } from '../../../styles';
import HeaderLogo from './HeaderLogo';
import HeaderNavigation from './HeaderNavigation';
import UserPersonalMenu from '../../user/UserPersonalMenu';
import ButtonThemeChange from '../../common/ButtonThemeChange';

interface HeaderWebProps {}

function HeaderWeb(props: HeaderWebProps) {
  return (
    <div css={block}>
      <HeaderLogo />
      <div css={itemBlock}>
        <HeaderNavigation />
        <UserPersonalMenu />
        <div css={themeChangeButtonBlock}>
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
  height: 3rem;
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

const themeChangeButtonBlock = css`
  padding-left: 0.5rem;
`;

export default HeaderWeb;
