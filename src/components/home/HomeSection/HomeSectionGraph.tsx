import { TestPage2 } from '../../../pages';
import homeSectionStyles from './homeSectionStyles';

interface HomeSectionGraphProps {}

function HomeSectionGraph(props: HomeSectionGraphProps) {
  return (
    <section css={homeSectionStyles.block}>
      <div css={homeSectionStyles.itemCommon}>
        <TestPage2 />
      </div>
    </section>
  );
}

export default HomeSectionGraph;
