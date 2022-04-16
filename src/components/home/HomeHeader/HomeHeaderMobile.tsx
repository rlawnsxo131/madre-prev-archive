import { css } from '@emotion/react';
import { media } from '../../../styles';
import ButtonThemeChange from '../../common/ButtonThemeChange';
import UserPersonalMenu from '../../user/UserPersonalMenu';
import HomeHeaderLogo from './HomeHeaderLogo';
import HomeHeaderMobileNavigation from './HomeHeaderMobileNavigation';

interface HomeHeaderMobileProps {}

function HomeHeaderMobile(props: HomeHeaderMobileProps) {
  return (
    <div css={block}>
      <HomeHeaderLogo />
      <div css={right}>
        <UserPersonalMenu />
        <ButtonThemeChange />
        <HomeHeaderMobileNavigation />
      </div>
    </div>
  );
}

const block = css`
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 3rem;

  ${media.xxxsmall} {
    width: 93%;
  }
  ${media.small} {
    display: none;
  }
`;

const right = css`
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1rem;
`;

export default HomeHeaderMobile;
