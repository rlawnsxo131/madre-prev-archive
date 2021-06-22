import { css } from '@emotion/react';
import media, { mediaQuery } from '../../styles/media';
import HomeHeader from './HomeHeader';

interface HomeTemplateProps {
  children: React.ReactNode;
}

function HomeTemplate({ children = null }: HomeTemplateProps) {
  return (
    <div css={block}>
      <HomeHeader />
      <main css={main}>
        <section css={section}>{children}</section>
      </main>
    </div>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
  align-items: center;
  border: 1px solid black;
`;

const main = css`
  border: 1px solid black;
  display: flex;
  justify-content: center;
  position: relative;
  padding-top: 4rem;
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

const section = css`
  border: 1px solid blue;
  width: 100%;
  ${media.xxxsmall} {
    width: 92%;
  }
  ${media.small} {
    width: 94%;
  }
`;

export default HomeTemplate;
