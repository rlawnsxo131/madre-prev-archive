import { css } from '@emotion/react';
import { NavLink } from 'react-router-dom';
import media, { mediaQuery } from '../../styles/media';
import palette from '../../styles/palette';
import zIndexes from '../../styles/zIndexes';

interface HomeHeaderProps {}

function HomeHeader(props: HomeHeaderProps) {
  return (
    <div css={block}>
      <div css={content}>
        <NavLink css={link} to="/">
          <h1>Data Visualizer</h1>
        </NavLink>
        <nav>navigations</nav>
      </div>
    </div>
  );
}

const block = css`
  position: fixed;
  display: flex;
  justify-content: center;
  align-items: center;
  top: 0;
  left: 0;
  width: 100%;
  z-index: ${zIndexes.homeHeader};
`;

const content = css`
  border: 1px solid red;
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
  ${mediaQuery(1300)} {
    width: 1260px;
  }
`;

const link = css`
  display: inline-flex;
  flex-flow: row wrap;
  font-size: 1.25rem;
  padding: 0.5rem 0.25rem 0.5rem 0.25rem;
  h1 {
    margin: 0;
    padding: 0;
    font-weight: bold;
    font-size: 1.5rem;
    color: ${palette.black};
  }
`;

export default HomeHeader;
