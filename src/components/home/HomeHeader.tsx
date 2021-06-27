import { css } from '@emotion/react';
import media, { mediaQuery } from '../../styles/media';
import zIndexes from '../../styles/zIndexes';
import HomeHeaderLogo from './HomeHeaderLogo';
import HomeNavigation from './HomeNavigation';

interface HomeHeaderProps {}

function HomeHeader(props: HomeHeaderProps) {
  return (
    <header css={block}>
      <div css={content}>
        <HomeHeaderLogo />
        <HomeNavigation />
      </div>
    </header>
  );
}

const block = css`
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0.25 1rem;
  z-index: ${zIndexes.homeHeader};
`;

const content = css`
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 5rem;
  ${media.xxxsmall} {
    width: 100%;
  }
  ${media.small} {
    width: 768px;
  }
  ${media.medium} {
    width: calc(100% - 4vw);
  }
  ${mediaQuery(1300)} {
    width: 1260px;
  }
`;

export default HomeHeader;
