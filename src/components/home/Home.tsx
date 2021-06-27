import { css } from '@emotion/react';
import HomeFirstSection from './HomeFirstSection';

interface HomeProps {}

function Home(props: HomeProps) {
  return (
    <div css={block}>
      <HomeFirstSection />
    </div>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
`;

export default Home;
