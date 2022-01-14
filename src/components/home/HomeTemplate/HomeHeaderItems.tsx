import { css } from '@emotion/react';
import HomeNavigation from './HomeNavigation';
import AuthGoogleButton from '../../auth/AuthGoogleButton/AuthGoogleButton';

interface HomeHeaderItemsProps {}

function HomeHeaderItems(props: HomeHeaderItemsProps) {
  return (
    <div css={block}>
      <div css={authBlock}>
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

const authBlock = css`
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0 0.5rem;
`;

export default HomeHeaderItems;
