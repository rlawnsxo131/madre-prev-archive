import { css } from '@emotion/react';
import MobilePublicNavigation from '../../navigation/MobilePublicNavigation';
import PublicNavigation from '../../navigation/PublicNavigation';
import UserPersonalMenu from '../../user/UserPersonalMenu';

interface HomeHeaderItemsProps {}

function HomeHeaderItems(props: HomeHeaderItemsProps) {
  return (
    <div css={block}>
      {/* navigation */}
      <PublicNavigation />
      <MobilePublicNavigation />

      {/* user menu */}
      <UserPersonalMenu />
    </div>
  );
}

const block = css`
  display: flex;
  align-items: center;
`;

export default HomeHeaderItems;
