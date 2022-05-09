import { css } from '@emotion/react';
import { palette, zIndexes } from '../../../styles';

type ToolTipPosition = 'top' | 'bottom';

interface ToolTipProps {
  children: React.ReactNode;
  position?: ToolTipPosition;
  visible: boolean;
  text: string;
}

function ToolTip({ children, position = 'top', visible, text }: ToolTipProps) {
  return (
    <div css={block}>
      {visible && (
        <div css={tooltip(position)}>
          <div css={triangle(position)} />
          <div css={textBlock}>{text}</div>
        </div>
      )}
      {children}
    </div>
  );
}

const block = css`
  position: relative;
  z-index: ${zIndexes.tooltip};
`;

const tooltip = (position: ToolTipPosition) => css`
  position: absolute;
  left: 0;
  display: flex;

  ${position === 'top' &&
  css`
    top: 0;
    flex-direction: column-reverse;
    transform: translateY(-100%);
  `}
  ${position === 'bottom' &&
  css`
    top: 100%;
    flex-direction: column;
  `};
`;

const triangle = (position: ToolTipPosition) => css`
  width: 0;
  height: 0;
  margin-left: 0.5rem;
  border: 6px solid transparent;

  ${position === 'top' &&
  css`
    border-top: 12px solid ${palette.opaque[50]};
    border-bottom: 0;
  `}

  ${position === 'bottom' &&
  css`
    border-top: 0;
    border-bottom: 12px solid ${palette.opaque[50]};
  `}
`;

const textBlock = css`
  padding: 0.5rem;
  border-radius: 3px;
  background: ${palette.opaque[50]};
`;

export default ToolTip;
