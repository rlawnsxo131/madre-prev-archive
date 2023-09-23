import { css } from '@emotion/react';
import useTransitionTimeoutEffect from '../../../hooks/useTransitionTimeoutEffect';
import { palette, themePalette, transitions, zIndexes } from '../../../styles';

type ScreenBaseType = 'default' | 'opaque';

interface ScreenBaseProps {
  children: React.ReactNode;
  visible: boolean;
  type?: ScreenBaseType;
}

function ScreenBase({ children, visible, type = 'default' }: ScreenBaseProps) {
  const closed = useTransitionTimeoutEffect({ visible });

  if (!visible && closed) return null;

  return <div css={block(visible, type)}>{children}</div>;
}

const block = (visible: boolean, type: ScreenBaseType) => css`
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: ${zIndexes.sliderBase};
  background: ${type === 'default' ? themePalette.bg4 : palette.opaque['50']};
  ${visible
    ? css`
        animation: ${transitions.slideUp} 0.5s forwards ease-in-out;
      `
    : css`
        animation: ${transitions.slideDown} 0.25s forwards ease-in-out;
      `};
`;

export default ScreenBase;
