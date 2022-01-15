import { css } from '@emotion/react';
import { navigationDisplay } from '../navigation.styles';
import PublicNavigationLinks from './PublicNavigationLinks';

interface PublicNavigationProps {}

function PublicNavigation(props: PublicNavigationProps) {
  return (
    <nav css={block}>
      <PublicNavigationLinks />
    </nav>
  );
}

const block = css`
  display: flex;
  justify-content: space-between;
  padding: 0 1rem;
  ${navigationDisplay};
`;

export default PublicNavigation;
