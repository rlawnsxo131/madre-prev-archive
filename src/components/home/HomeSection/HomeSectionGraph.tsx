import homeSectionStyles from './homeSectionStyles';

interface HomeSectionGraphProps {}

function HomeSectionGraph(props: HomeSectionGraphProps) {
  return (
    <section css={homeSectionStyles.block}>
      <div css={homeSectionStyles.itemCommon}>chart area</div>
    </section>
  );
}

export default HomeSectionGraph;
