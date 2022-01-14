import { RealTimeAsyncImage } from '../../../image/images';
import HomeSectionStyles from './HomeSection.styles';

interface HomeSectionThinkAboutProps {}

function HomeSectionThinkAbout(props: HomeSectionThinkAboutProps) {
  return (
    <section css={HomeSectionStyles.section}>
      <div css={HomeSectionStyles.descriptionBlock}>
        <div css={HomeSectionStyles.description}>
          <h3>Think About</h3>
          <h2>우린 지금, 얼마나 많은 데이터와 함께할까요?</h2>
          <p>
            그리고 어떤 의미가 될 수 있을까요.{`\n`}
            어쩌면 우리가 생각하는 그 이상으로 많은 곳에서 함께 하고 있었을
            거예요.
          </p>
        </div>
      </div>
      <div css={HomeSectionStyles.imageBlock}>
        <RealTimeAsyncImage />
      </div>
    </section>
  );
}

export default HomeSectionThinkAbout;
