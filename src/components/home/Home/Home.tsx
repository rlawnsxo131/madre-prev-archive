import { css } from '@emotion/react';
import HomeSectionThinkAbout from './HomeSectionThinkAbout';

interface HomeProps {}

function Home(props: HomeProps) {
  return (
    <div css={block}>
      <HomeSectionThinkAbout />
      <HomeSectionThinkAbout />
    </div>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
`;

export default Home;
