import { css } from '@emotion/react';

interface HomeSectionSecondProps {}

function HomeSectionSecond(props: HomeSectionSecondProps) {
  return <section css={block}>second section</section>;
}

const block = css`
  display: flex;
  justify-content: center;
  flex-flow: row wrap;
`;

export default HomeSectionSecond;
