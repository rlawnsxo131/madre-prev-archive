import { css } from '@emotion/react';
import InsertImage from '../../../image/images/InsertImage';
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
          <p css={homeP}>그리고 우리에게 어떤 의미가 될 수 있을까요.</p>
        </div>
      </div>
      <div css={imageBlock}>
        <InsertImage />
      </div>
    </section>
  );
}

const descriptionBlock = css`
  ${homeBlockItemCommon};
  flex: 1;
  display: flex;
  align-items: center;
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
    width: 100%;
    height: auto;
  }
`;

export default HomeSectionThinkAbout;
