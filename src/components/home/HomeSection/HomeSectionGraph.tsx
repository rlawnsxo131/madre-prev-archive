import TestPage from '../../../pages/TestPage';
import homeSectionStyles from './homeSectionStyles';

interface HomeSectionGraphProps {}

function HomeSectionGraph(props: HomeSectionGraphProps) {
  return (
    <section css={homeSectionStyles.block}>
      <div css={homeSectionStyles.itemCommon}>
        <TestPage />
      </div>
    </section>
  );
}

export default HomeSectionGraph;
