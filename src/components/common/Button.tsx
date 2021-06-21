import { css } from '@emotion/react';

interface ButtonProps {}

function Button(props: ButtonProps) {
  return <button css={block}>button</button>;
}

const block = css``;

export default Button;
