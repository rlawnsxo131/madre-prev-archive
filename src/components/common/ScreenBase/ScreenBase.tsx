import { css } from '@emotion/react';
import useThemeState from '../../../hooks/theme/useThemeState';
import useTransitionTimeoutEffect from '../../../hooks/useTransitionTimeoutEffect';
import { Theme } from '../../../redux/theme';
import { palette, themeColor, transitions, zIndexes } from '../../../styles';

interface ScreenBaseProps {
  children: React.ReactNode;
  visible: boolean;
}

function ScreenBase({ children, visible }: ScreenBaseProps) {
  const { theme } = useThemeState();
  const closed = useTransitionTimeoutEffect({ visible });

  if (!visible && closed) return null;

  return <div css={block(theme, visible)}>{children}</div>;
}

const block = (theme: Theme, visible: boolean) => css`
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: ${zIndexes.sliderBase};
  background: ${palette.opaque[50]};
  ${visible
    ? css`
        animation: ${transitions.slideUp} 0.4s forwards ease-in-out;
      `
    : css`
        animation: ${transitions.slideDown} 0.25s forwards ease-in-out;
      `};
`;

export default ScreenBase;
