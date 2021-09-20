import { css } from '@emotion/react';
import { useEffect, useState } from 'react';
import {
  DarkmodeThemeType,
  useDarkmodeValue,
} from '../../../atoms/darkmodeState';
import { themeColor } from '../../../styles/palette';
import transitions from '../../../styles/transitions';
import OpaqueLayer from '../OpaqueLayer';

interface PopupBaseProps {
  children: React.ReactNode;
  visible: boolean;
}

function PopupBase({ children, visible }: PopupBaseProps) {
  const { theme } = useDarkmodeValue();
  const [closed, setClosed] = useState(true);

  useEffect(() => {
    let timeoutId: NodeJS.Timeout | null = null;
    if (visible) {
      setClosed(false);
    } else {
      timeoutId = setTimeout(() => {
        setClosed(true);
      }, 250);
    }
    return () => {
      if (timeoutId) {
        clearTimeout(timeoutId);
      }
    };
  }, [visible]);

  if (!visible && closed) return null;

  return (
    <OpaqueLayer visible={visible}>
      <div css={popup(visible, theme)}>{children}</div>
    </OpaqueLayer>
  );
}

const popup = (visible: boolean, theme: DarkmodeThemeType) => css`
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
