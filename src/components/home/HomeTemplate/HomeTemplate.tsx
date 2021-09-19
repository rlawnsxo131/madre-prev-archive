import { css } from '@emotion/react';
import media, { mediaQuery } from '../../../styles/media';
import HomeHeader from './HomeHeader';

interface HomeTemplateProps {
  children: React.ReactNode;
}

function HomeTemplate({ children }: HomeTemplateProps) {
  return (
    <div css={block}>
      <HomeHeader />
      <main css={main}>
        <div css={content}>{children}</div>
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
  padding: 4rem 0.5rem 0 0.5rem;
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

const content = css`
  display: flex;
  flex-direction: column;
  width: 100%;
`;

export default HomeTemplate;
