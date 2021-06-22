import { css } from '@emotion/react';
import HomeTemplate from './HomeTemplate';

interface HomeProps {}

function Home(props: HomeProps) {
  return (
    <HomeTemplate>
      <div css={block}>home main content</div>
    </HomeTemplate>
  );
}

const block = css``;

export default Home;
