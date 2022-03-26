import { css } from '@emotion/react';
import { palette, themePalette } from '../../../styles';
import { InputSize, inputSizeMap } from './Input.styles';

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
  width: ${inputSizeMap[size].width};
  padding: ${inputSizeMap[size].padding};
  background: inherit;
  border: 2px solid ${themePalette.border_element1};
  border-radius: 3px;
  caret-color: ${palette.pink['500']};

  &:active,
  &:focus {
    border: 2px solid ${palette.pink['500']};
  }
`;

export default Input;
