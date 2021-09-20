import { css } from '@emotion/react';
import { memo } from 'react';
import {
  DarkmodeThemeType,
  useDarkmodeValue,
} from '../../../atoms/darkmodeState';
import {
  buttonColorMap,
  ButtonColorType,
  ButtonShapeType,
  buttonSizeMap,
  ButtonSizeType,
  outlineButtonStyle,
} from '../../../styles/buttonStyle';
import palette from '../../../styles/palette';

interface ButtonProps extends Omit<React.HTMLProps<HTMLButtonElement>, 'size'> {
  children: React.ReactNode;
  color?: ButtonColorType;
  size?: ButtonSizeType;
  shape?: ButtonShapeType;
  outline?: boolean;
  icon?: React.ReactNode;
}
function Button({
  children,
  color = 'gray',
  size = 'medium',
  shape = 'rect',
  outline = false,
  icon,
  ...rest
}: ButtonProps) {
  const htmlProps = rest as any;
  const { theme } = useDarkmodeValue();
  return (
    <button
      css={block(color, size, shape, outline, theme)}
      onClick={(e) => {
        if (htmlProps.onClick) {
          htmlProps.onClick(e);
        }
        (e.target as HTMLButtonElement).blur();
      }}
      {...htmlProps}
    >
      {icon && <div css={iconStyle}>{icon}</div>}
      {children}
    </button>
  );
}

const block = (
  color: ButtonColorType,
  size: ButtonSizeType,
  shape: ButtonShapeType,
  outline: boolean,
  theme: DarkmodeThemeType,
) => css`
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  cursor: pointer;
  outline: none;
  border: none;
  box-sizing: border-box;
  cursor: pointer;
  border-radius: 4px;
  border: 1px solid ${buttonColorMap[color].background};

  ${outline &&
  css`
    color: ${buttonColorMap[color].background};
    background: ${outlineButtonStyle[theme].background};
    &:hover {
      color: ${buttonColorMap[color].hoverBackground};
      border: 1px solid ${buttonColorMap[color].hoverBackground};
      opacity: ${outlineButtonStyle[theme].opacity};
    }
  `}
  ${!outline &&
  css`
    color: ${palette.white};
    background: ${buttonColorMap[color].background};
    &:hover {
      background: ${buttonColorMap[color].hoverBackground};
      border: 1px solid ${buttonColorMap[color].hoverBackground};
    }
  `}
  
  &:disabled {
    cursor: not-allowed;
    border: 1px solid ${palette.gray['100']};
    color: ${palette.gray['500']};
    ${outline &&
    css`
      background: ${outlineButtonStyle[theme].background};
      &:hover {
        background: ${outlineButtonStyle[theme].background};
      }
    `}
    ${!outline &&
    css`
      background: ${palette.gray['300']};
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
`;

export default memo(Button);
