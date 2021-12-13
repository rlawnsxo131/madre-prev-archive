import { css } from '@emotion/react';
import { RealTimeAsyncImage } from '../../../image/images';
import homeStyles from './homeStyles';

interface HomeSectionThinkAboutProps {}

function HomeSectionThinkAbout(props: HomeSectionThinkAboutProps) {
  return (
    <section css={homeStyles.block}>
      <div css={descriptionBlock}>
        <div css={description}>
          <h3>Think About</h3>
          <h2>우린 지금, 얼마나 많은 데이터와 함께할까요?</h2>
          <p>
            그리고 어떤 의미가 될 수 있을까요.{`\n`}
            어쩌면 우리가 생각하는 그 이상으로 많은 곳에서 함께 하고 있었을
            거예요.
          </p>
        </div>
      </div>
      <div css={imageBlock}>
        <RealTimeAsyncImage />
      </div>
    </section>
  );
}

const descriptionBlock = css`
  ${homeStyles.block};
  flex: 1;
  display: flex;
`;

const description = css`
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  gap: 1.25rem;
`;

const imageBlock = css`
  ${homeStyles.itemCommon};
  flex: 1;
  display: flex;
  justify-content: center;
  gap: 1.25rem;
  svg {
    max-width: 35rem;
    height: auto;
  }
`;

export default HomeSectionThinkAbout;
