import type { PropsWithChildren } from 'react';
import { useEffect, useRef, useState } from 'react';

export type OpaqueLayerProps = PropsWithChildren<{
  visible: boolean;
  delay?: number;
}>;

export function OpaqueLayer({
  children,
  visible,
  delay = 250,
}: OpaqueLayerProps) {
  const [closed, setClosed] = useState(true);
  const [animate, setAnimate] = useState(false);
  const mounted = useRef(false);
  const timeoutId = useRef<NodeJS.Timeout | null>(null);

  useEffect(() => {
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
      }, delay);
    }

    if (visible) {
      setClosed(false);
    }

    return () => {
      document.body.style.overflowY = 'initial';

      if (timeoutId.current) {
        clearTimeout(timeoutId.current);
      }
    };
  }, [visible, delay]);

  if (!animate && !visible && closed) return null;

  return <div>{children}</div>;
}
