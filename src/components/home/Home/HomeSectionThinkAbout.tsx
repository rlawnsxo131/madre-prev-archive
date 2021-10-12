import { css } from '@emotion/react';
import { LostOnlineImage } from '../../../image/images';
import {
  homeBlock,
  homeBlockItemCommon,
  homeH3,
  homeH5,
  homeP,
} from './homeStyles';

interface HomeSectionThinkAboutProps {}

function HomeSectionThinkAbout(props: HomeSectionThinkAboutProps) {
  return (
    <section css={homeBlock}>
      <div css={descriptionBlock}>
        <div css={description}>
          <h5 css={homeH5}>Think About</h5>
          <h3 css={homeH3}>우린 지금, 얼마나{'\n'}많은 데이터와 함께할까요?</h3>
          <p css={homeP}>
            어쩌면, 저 밤하늘에{'\n'}반짝이는 것들만큼이나 많을지도 몰라요.
          </p>
        </div>
      </div>
      <div css={imageBlock}>
        <LostOnlineImage />
      </div>
    </section>
  );
}

const descriptionBlock = css`
  ${homeBlockItemCommon};
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
`;

const description = css`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  gap: 1.25rem;
`;

const imageBlock = css`
  ${homeBlockItemCommon};
  flex: 1;
  display: flex;
  gap: 1.25rem;
  svg {
    height: auto;
    width: 100%;
  }
`;

export default HomeSectionThinkAbout;
