import { css } from '@emotion/react';
import { InputSize } from './Input.styles';

interface InputProps
  extends Omit<React.InputHTMLAttributes<HTMLInputElement>, 'size'> {
  size?: InputSize;
}

function Input({ size = 'medium', ...props }: InputProps) {
  return <input css={block(size)} {...props} />;
}

const block = (size: InputSize) => css``;

export default Input;
