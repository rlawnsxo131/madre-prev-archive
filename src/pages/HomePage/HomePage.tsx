import { css } from '@emotion/react';
import {
  HomeSectionGraph,
  HomeSectionThinkAbout,
} from '../../components/home/HomeSection';

interface HomePageProps {}

function HomePage(props: HomePageProps) {
  return (
    <div css={block}>
      <HomeSectionThinkAbout />
      <HomeSectionGraph />
    </div>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
`;

export default HomePage;
