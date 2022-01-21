import { css } from '@emotion/react';
import HomeHeaderMobileNavigationButton from './HomeHeaderMobileNavigationButton';
import HomeHeaderMobileNavigationLinks from './HomeHeaderMobileNavigationLinks';

interface HomeHeaderMobileNavigationProps {}

function HomeHeaderMobileNavigation(props: HomeHeaderMobileNavigationProps) {
  return (
    <div css={block}>
      <HomeHeaderMobileNavigationButton />
      <HomeHeaderMobileNavigationLinks />
    </div>
  );
}

const block = css`
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  &:nth-of-type(1) {
    padding: 0 0.5rem;
  }
  &:nth-of-type(2) {
    padding-left: 0.5rem;
  }
`;

export default HomeHeaderMobileNavigation;
