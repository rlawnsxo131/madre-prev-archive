import { css } from '@emotion/react';
import PublicNavigation from '../../navigation/PublicNavigation';
import UserPersonalMenu from '../../user/UserPersonalMenu';

interface HomeHeaderItemsProps {}

function HomeHeaderItems(props: HomeHeaderItemsProps) {
  return (
    <div css={block}>
      <PublicNavigation />
      <UserPersonalMenu />
    </div>
  );
}

const block = css`
  display: flex;
  align-items: center;
`;

export default HomeHeaderItems;
