import { css } from '@emotion/react';
import { palette } from '../../../styles';
import { InputSize } from './Input.styles';

interface InputProps
  extends Omit<React.InputHTMLAttributes<HTMLInputElement>, 'size'> {
  size?: InputSize;
}

function Input({ size = 'medium', ...props }: InputProps) {
  return <input css={block(size)} {...props} />;
}

const block = (size: InputSize) => css`
  all: unset;
  outline: none;
  padding: 0.3rem;
  background: inherit;
  border: 1px solid ${palette.pink['500']};
  border-radius: 3px;
  caret-color: ${palette.pink['500']};

  &:active,
  &:focus {
    border: 2px solid ${palette.pink['500']};
  }

  ${size === 'small' &&
  css`
    width: 7.5rem;
  `}

  ${size === 'medium' &&
  css`
    width: 18.75rem;
  `}

  ${size === 'responsive' &&
  css`
    width: 100%;
  `}
`;

export default Input;
