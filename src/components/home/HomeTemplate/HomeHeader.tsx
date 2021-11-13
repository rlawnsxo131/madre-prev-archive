import { css } from '@emotion/react';
import zIndexes from '../../../styles/zIndexes';
import HomeHeaderLogo from './HomeHeaderLogo';
import HomeHeaderItems from './HomeHeaderItems';
import { media, mediaQuery } from '../../../styles';

interface HomeHeaderProps {}

function HomeHeader(props: HomeHeaderProps) {
  return (
    <header css={block}>
      <div css={content}>
        <HomeHeaderLogo />
        <HomeHeaderItems />
      </div>
    </header>
  );
}

const block = css`
  position: sticky;
  top: 0;
  left: 0;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0.25rem 0;
  z-index: ${zIndexes.homeHeader};
`;

const content = css`
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 4rem;
  ${media.xxxsmall} {
    width: 100%;
  }
  ${media.small} {
    width: 768px;
  }
  ${media.medium} {
    width: calc(100% - 4vw);
  }
  ${mediaQuery(1285)} {
    width: 1250px;
  }
`;

export default HomeHeader;
