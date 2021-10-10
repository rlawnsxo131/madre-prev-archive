import { css } from '@emotion/react';
import HomeSectionGraph from './HomeSectionGraph';
import HomeSectionThinkAbout from './HomeSectionThinkAbout';

interface HomeProps {}

function Home(props: HomeProps) {
  return (
    <div css={block}>
      <HomeSectionGraph />
      <HomeSectionThinkAbout />
    </div>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
`;

export default Home;
