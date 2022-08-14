// import { TestPage2 } from '../../../pages';
import MadreMadreStoryStyles from './MadreStory.styles';

interface MadreStorySectionGraphProps {}

function MadreStorySectionGraph(props: MadreStorySectionGraphProps) {
  return (
    <section css={MadreMadreStoryStyles.section}>
      <div css={MadreMadreStoryStyles.itemCommon}>{/* <TestPage2 /> */}</div>
    </section>
  );
}

export default MadreStorySectionGraph;
