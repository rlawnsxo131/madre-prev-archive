import { isClient } from '@madre/utils';
import { useEffect, useLayoutEffect } from 'react';

/**
 * @description client 에서는 useLayoutEffect 방식을 쓰고, server 에서는 useEffect 방식을 쓰도록 합니다.
 * server 에서는 useLayoutEffect 함수를 사용하면 warning 오류가 발생하기 때문에 사용하는 함수입니다.
 */
export const useIsomorphicLayoutEffect = isClient()
  ? useLayoutEffect
  : useEffect;
