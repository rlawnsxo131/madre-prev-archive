import { css } from '@emotion/react';

interface HomeSectionSecondProps {}

function HomeSectionSecond(props: HomeSectionSecondProps) {
  return (
    <section css={block}>
      second
    </section>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

export default HomeSectionSecond;
