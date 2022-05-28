import { useState, useEffect } from 'react';

interface UseTransitionTimeoutEffectParams {
  visible: boolean;
  delay?: number;
}

export default function useTransitionTimeoutEffect({
  visible,
  delay = 250,
}: UseTransitionTimeoutEffectParams) {
  const [closed, setClosed] = useState(true);

  useEffect(() => {
    let timeoutId: NodeJS.Timeout | null = null;

    if (visible) {
      setClosed(false);
    } else {
      timeoutId = setTimeout(() => {
        setClosed(true);
      }, delay);
    }

    return () => {
      if (timeoutId) {
        clearTimeout(timeoutId);
      }
    };
  }, [visible, delay]);

  return closed;
}
