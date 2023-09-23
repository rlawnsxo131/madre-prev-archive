import { isClient } from '@madre/common';
import { useEffect, useLayoutEffect } from 'react';

export const useIsomorphicLayoutEffect = isClient()
  ? useLayoutEffect
  : useEffect;
