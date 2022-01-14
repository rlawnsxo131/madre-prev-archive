import { css } from '@emotion/react';
import HomeNavigation from './HomeNavigation';
import AuthGoogleButton from '../../auth/AuthGoogleButton/AuthGoogleButton';
import homeTemplateStyles from './homeTemplateStyles';

interface HomeHeaderItemsProps {}

function HomeHeaderItems(props: HomeHeaderItemsProps) {
  return (
    <div css={block}>
      <div css={homeTemplateStyles.itemBlock}>
        <AuthGoogleButton />
      </div>
      <HomeNavigation />
    </div>
  );
}

const block = css`
  position: relative;
  display: flex;
  align-items: center;
`;

export default HomeHeaderItems;
