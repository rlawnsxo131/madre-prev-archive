import { css } from '@emotion/react';
import { CountingStarsImage, WorkingLateImage } from '../../../image/images';
import media from '../../../styles/media';
import { themeColor } from '../../../styles/palette';
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
            어쩌면, 밤하늘에{'\n'}빛나는 별만큼이나 많을지도 몰라요.
          </p>
        </div>
      </div>
      <div css={imageBlock}>
        <CountingStarsImage />
        <WorkingLateImage />
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
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 1.25rem;
`;

const imageBlock = css`
  ${homeBlockItemCommon};
  flex: 1;
  display: flex;
  flex-flow: row wrap;
  justify-content: space-around;
  gap: 1.25rem;
  background: ${themeColor.content['light']};
  svg {
    height: auto;
    ${media.xxxsmall} {
      width: 100%;
    }
    ${media.xxsmall} {
      width: 15rem;
    }
  }
`;

export default HomeSectionThinkAbout;
