import { css } from '@emotion/react';
import { forwardRef, memo } from 'react';
import { palette, themePalette } from '../../../styles';
import { InputSize, inputSizeMap } from './Input.styles';

interface InputProps
  extends Omit<React.InputHTMLAttributes<HTMLInputElement>, 'size'> {
  size?: InputSize;
}

const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ size = 'medium', ...props }, ref) => {
    return <input css={input(size)} {...props} ref={ref} />;
  },
);
Input.displayName = 'Input';

const input = (size: InputSize) => css`
  margin: 0;
  font: inherit;
  color: ${themePalette.text1};
  outline: none;
  width: ${inputSizeMap[size].width};
  padding: ${inputSizeMap[size].padding};
  background: inherit;
  border: 2px solid ${themePalette.border_element1};
  border-radius: 3px;
  caret-color: ${palette.blue['500']};

  &:active,
  &:focus {
    border: 2px solid ${palette.blue['500']};
  }

  &:-webkit-autofill,
  &:-webkit-autofill:hover,
  &:-webkit-autofill:focus,
  &:-webkit-autofill:active {
    transition: background-color 5000s;
    -webkit-text-font-size: inherit !important;
    -webkit-text-fill-color: ${themePalette.text1} !important;
  }
`;

export default memo(Input);
