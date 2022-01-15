import { css } from '@emotion/react';
import MobilePublicNavigation from '../../navigation/MobilePublicNavigation';
import PublicNavigation from '../../navigation/PublicNavigation';
import UserPersonalMenu from '../../user/UserPersonalMenu';

interface HomeHeaderItemsProps {}

function HomeHeaderItems(props: HomeHeaderItemsProps) {
  return (
    <div css={block}>
      <PublicNavigation />
      <UserPersonalMenu />

      <MobilePublicNavigation />
    </div>
  );
}

const block = css`
  display: flex;
  align-items: center;
`;

export default HomeHeaderItems;
