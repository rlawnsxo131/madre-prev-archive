import { css } from '@emotion/react';

interface ToolTipProps {}

function ToolTip(props: ToolTipProps) {
  return <div css={block}></div>;
}

const block = css`
  position: relative;
`;

export default ToolTip;
