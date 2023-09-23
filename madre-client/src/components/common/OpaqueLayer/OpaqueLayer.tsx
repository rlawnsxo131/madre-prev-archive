import { useEffect, useRef, useState } from 'react';
import { css } from '@emotion/react';
import { palette, transitions, zIndexes } from '../../../styles';

interface OpaqueLayerProps {
  visible: boolean;
}

function OpaqueLayer({ visible }: OpaqueLayerProps) {
  const [closed, setClosed] = useState(true);
  const [animate, setAnimate] = useState(false);
  const mounted = useRef(false);
  const timeoutId = useRef<NodeJS.Timeout | null>(null);

  useEffect(() => {
    // scrollbar
    document.body.style.overflowY = visible ? 'hidden' : 'initial';

    if (!mounted.current) {
      mounted.current = true;
    } else {
      setAnimate(true);
      timeoutId.current = setTimeout(() => {
        setAnimate(false);
        if (!visible) {
          setClosed(true);
        }
      }, 250);
    }

    if (visible) {
      setClosed(false);
    }

    return () => {
      if (timeoutId.current) {
        clearTimeout(timeoutId.current);
      }
    };
  }, [visible]);

  useEffect(() => {
    return () => {
      document.body.style.overflowY = 'initial';
    };
  }, []);

  if (!animate && !visible && closed) return null;

  return <div css={block(visible)} />;
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
  background: ${palette.opaque[50]};
  ${visible
    ? css`
        animation: ${transitions.fadeIn} 0.25s forwards;
      `
    : css`
        animation: ${transitions.fadeOut} 0.25s forwards;
      `}
`;

export default OpaqueLayer;
