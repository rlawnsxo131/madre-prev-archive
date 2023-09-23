import { isClient } from '@madre/utils';
import { useEffect, useLayoutEffect } from 'react';

export const useIsomorphicLayoutEffect = isClient()
  ? useLayoutEffect
  : useEffect;
