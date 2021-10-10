import TestPage from '../../../pages/TestPage';
import { homeBlock, homeBlockItemCommon } from './homeStyles';

interface HomeSectionGraphProps {}

function HomeSectionGraph(props: HomeSectionGraphProps) {
  return (
    <section css={homeBlock}>
      <div css={homeBlockItemCommon}>
        <TestPage />
      </div>
    </section>
  );
}

export default HomeSectionGraph;
