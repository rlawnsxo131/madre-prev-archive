import { css } from '@emotion/react';
import useTransitionTimeoutEffect from '../../../hooks/useTransitionTimeoutEffect';
import { transitions, zIndexes } from '../../../styles';
import { themePalette } from '../../../styles';
import OpaqueLayer from '../OpaqueLayer';

interface PopupBaseProps {
  children: React.ReactNode;
  visible: boolean;
}

function PopupBase({ children, visible }: PopupBaseProps) {
  const closed = useTransitionTimeoutEffect({ visible });

  if (!visible && closed) return null;

  return (
    <>
      <OpaqueLayer visible={visible} />
      <div css={block}>
        <div css={content(visible)}>{children}</div>
      </div>
    </>
  );
}

const block = css`
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: ${zIndexes.popup};
  display: flex;
  align-items: center;
  justify-content: center;
`;

const content = (visible: boolean) => css`
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 0.25rem;
  background: ${themePalette.bg_element1};
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.09);
  z-index: ${zIndexes.popup};
  ${visible
    ? css`
        animation: ${transitions.popInFromBottom} 0.4s forwards ease-in-out;
      `
    : css`
        animation: ${transitions.popOutToBottom} 0.25s forwards ease-in-out;
      `};
`;

export default PopupBase;
