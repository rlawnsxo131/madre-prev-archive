import { css } from '@emotion/react';
import { ColorTheme, useColorThemeValue } from '../../../atoms/colorThemeState';
import useTransitionTimeoutEffect from '../../../lib/hooks/useTransitionTimeoutEffect';
import { themeColor } from '../../../styles/palette';
import transitions from '../../../styles/transitions';
import OpaqueLayer from '../OpaqueLayer';

interface PopupBaseProps {
  children: React.ReactNode;
  visible: boolean;
}

function PopupBase({ children, visible }: PopupBaseProps) {
  const { theme } = useColorThemeValue();
  const closed = useTransitionTimeoutEffect({ visible });

  if (!visible && closed) return null;

  return (
    <OpaqueLayer visible={visible}>
      <div css={popup(visible, theme)}>{children}</div>
    </OpaqueLayer>
  );
}

const popup = (visible: boolean, theme: ColorTheme) => css`
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 3px;
  background: ${themeColor.background[theme]};
  ${visible
    ? css`
        animation: ${transitions.popInFromBottom} 0.25s forwards ease-in-out;
      `
    : css`
        animation: ${transitions.popOutToBottom} 0.25s forwards ease-in-out;
      `};
`;

export default PopupBase;
