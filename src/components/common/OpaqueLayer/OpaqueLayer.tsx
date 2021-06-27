import { css } from '@emotion/react';
import transitions from '../../../styles/transitions';
import zIndexes from '../../../styles/zIndexes';

interface OpaqueLayerProps {
  children: React.ReactNode;
  visible: boolean;
}

function OpaqueLayer({ visible, children }: OpaqueLayerProps) {
  return <div css={block(visible)}>{children}</div>;
}

const block = (visible: boolean) => css`
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: ${zIndexes.opaqueLayer};
  background: rgba(100, 100, 100, 0.5);
  ${visible
    ? css`
        animation: ${transitions.fadeIn} 0.25s forwards;
      `
    : css`
        animation: ${transitions.fadeOut} 0.25s forwards;
      `}
`;

export default OpaqueLayer;
