import { css } from '@emotion/react';
import { memo } from 'react';
import { palette, basicStyles } from '../../../styles';
import {
  ButtonColor,
  ButtonShape,
  ButtonSize,
  buttonSizeMap,
  buttonColorMap,
  buttonOutlineStyle,
} from './Button.styles';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
  color?: ButtonColor;
  size?: ButtonSize;
  shape?: ButtonShape;
  outline?: boolean;
  icon?: React.ReactNode;
  buttonRef?: React.RefObject<HTMLButtonElement>;
}

function Button({
  children,
  color = 'gray',
  size = 'medium',
  shape = 'rect',
  outline = false,
  icon,
  buttonRef,
  ...rest
}: ButtonProps) {
  return (
    <button
      ref={buttonRef}
      css={[basicStyles.button, button(color, size, shape, outline)]}
      onClick={(e) => {
        if (rest.onClick) {
          rest.onClick(e);
        }
        e.currentTarget.blur();
      }}
      {...rest}
    >
      {icon && <div css={iconStyle}>{icon}</div>}
      {children}
    </button>
  );
}

const button = (
  color: ButtonColor,
  size: ButtonSize,
  shape: ButtonShape,
  outline: boolean,
) => css`
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font-size: 1rem;
  font-weight: 600;

  border-radius: 4px;
  border: 1px solid ${buttonColorMap[color].background};

  ${outline &&
  css`
    color: ${buttonOutlineStyle[color].default};
    background: ${buttonOutlineStyle.background};
    border: 1px solid ${buttonOutlineStyle[color].default};
    &:hover {
      color: ${buttonOutlineStyle[color].hover};
      border: 1px solid ${buttonOutlineStyle[color].hover};
      opacity: ${buttonOutlineStyle.opacity};
    }
  `}
  ${!outline &&
  css`
    color: ${buttonColorMap[color].font};
    background: ${buttonColorMap[color].background};
    &:hover {
      background: ${buttonColorMap[color].hoverBackground};
      border: 1px solid ${buttonColorMap[color].hoverBackground};
    }
  `}
  
  &:disabled {
    cursor: not-allowed;
    ${outline &&
    css`
      color: ${buttonOutlineStyle.disabled};
      background: ${buttonOutlineStyle.background};
      border: 1px solid ${buttonOutlineStyle.disabled};
      &:hover {
        background: ${buttonOutlineStyle.background};
        opacity: 1;
      }
    `}
    ${!outline &&
    css`
      color: ${palette.gray['500']};
      background: ${palette.gray['300']};
      border: 1px solid ${palette.gray['100']};
      &:hover {
        background: ${palette.gray['300']};
      }
    `}
  }

  ${shape === 'round' &&
  css`
    border-radius: 2rem;
  `}

  ${size === 'responsive' &&
  css`
    flex: 1;
    width: 100%;
    min-height: 2.5rem;
    height: auto;
    font-size: 1.125rem;
  `};

  ${size !== 'responsive' &&
  css`
    height: ${buttonSizeMap[size].height};
    padding: ${buttonSizeMap[size].padding};
    font-size: ${buttonSizeMap[size].fontSize};
  `};
`;

const iconStyle = css`
  display: flex;
  justify-content: center;
  align-items: center;
  margin-right: 0.4rem;
  img,
  svg {
    width: 1rem;
    height: 1rem;
  }
`;

export default memo(Button);
