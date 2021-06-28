import { css } from '@emotion/react';
import HomeSectionFirst from './HomeSectionFirst';
import HomeSectionSecond from './HomeSectionSecond';

interface HomeProps {}

function Home(props: HomeProps) {
  return (
    <div css={block}>
      <HomeSectionFirst />
      <HomeSectionSecond />
    </div>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
`;

export default Home;
