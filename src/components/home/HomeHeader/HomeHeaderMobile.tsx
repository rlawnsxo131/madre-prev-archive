import { css } from '@emotion/react';
import { media } from '../../../styles';
import UserPersonalMenu from '../../user/UserPersonalMenu';
import homeStyles from '../home.styles';
import HomeHeaderStyles from './HomeHeader.styles';
import HomeHeaderLogo from './HomeHeaderLogo';
import HomeHeaderMobileNavigation from './HomeHeaderMobileNavigation';

interface HomeHeaderMobileProps {}

function HomeHeaderMobile(props: HomeHeaderMobileProps) {
  return (
    <div css={block}>
      <HomeHeaderMobileNavigation />
      <HomeHeaderLogo />
      <UserPersonalMenu />
    </div>
  );
}

const block = css`
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 3rem;
  div:nth-of-type(1) {
    position: absolute;
    left: 0;
  }
  div:nth-of-type(3) {
    position: absolute;
    right: 0;
  }
  /* div:nth-of-type(2) {
    transform: translateX(25%);
  } */
  ${media.xxxsmall} {
    width: 93%;
  }
  ${media.small} {
    display: none;
  }
`;

export default HomeHeaderMobile;
