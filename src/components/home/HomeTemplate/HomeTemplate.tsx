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
  padding: 0 0.5rem;
`;

const main = css`
  display: flex;
  justify-content: center;
  position: relative;
  ${media.xxxsmall} {
    width: 100%;
  }
  ${mediaQuery(800)} {
    width: 768px;
  }
  ${media.medium} {
    width: 96%;
  }
  ${mediaQuery(1285)} {
    width: 1250px;
  }
`;

const content = css`
  display: flex;
  flex-direction: column;
  width: 100%;
`;

export default HomeTemplate;
