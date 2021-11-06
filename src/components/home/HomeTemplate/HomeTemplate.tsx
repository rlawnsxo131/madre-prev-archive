import { css } from '@emotion/react';
import { Outlet } from 'react-router-dom';
import { media, mediaQuery } from '../../../styles';
import HomeHeader from './HomeHeader';

interface HomeTemplateProps {}

function HomeTemplate(props: HomeTemplateProps) {
  return (
    <div css={block}>
      <HomeHeader />
      <main css={main}>
        <div css={content}>
          <Outlet />
        </div>
      </main>
    </div>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const main = css`
  display: flex;
  justify-content: center;
  position: relative;
  padding: 0 0.5rem;
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
    width: 1250px;
  }
`;

const content = css`
  display: flex;
  flex-direction: column;
  width: 100%;
`;

export default HomeTemplate;
