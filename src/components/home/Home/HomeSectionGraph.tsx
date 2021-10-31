import TestPage from '../../../pages/TestPage';
import homeStyles from './homeStyles';

interface HomeSectionGraphProps {}

function HomeSectionGraph(props: HomeSectionGraphProps) {
  return (
    <section css={homeStyles.block}>
      <div css={homeStyles.itemCommon}>
        <TestPage />
      </div>
    </section>
  );
}

export default HomeSectionGraph;
