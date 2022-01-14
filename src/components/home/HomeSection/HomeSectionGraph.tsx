import { TestPage2 } from '../../../pages';
import HomeSectionStyles from './HomeSection.styles';

interface HomeSectionGraphProps {}

function HomeSectionGraph(props: HomeSectionGraphProps) {
  return (
    <section css={HomeSectionStyles.section}>
      <div css={HomeSectionStyles.itemCommon}>
        <TestPage2 />
      </div>
    </section>
  );
}

export default HomeSectionGraph;
