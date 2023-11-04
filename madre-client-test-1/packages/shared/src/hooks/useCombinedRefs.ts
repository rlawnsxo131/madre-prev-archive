import type { MutableRefObject, Ref } from 'react';
import { useCallback } from 'react';

type CallbackRef<T> = (ref: T | null) => void;

/**
 *
 * @description 여러개의 ref 를 하나로 합치고 싶을때 사용합니다.
 *
 * @param Array<Ref<T> | CallbackRef<T>>
 * @returns Ref<T>
 */
export function useCombinedRefs<T>(
  ...refs: Array<Ref<T> | CallbackRef<T>>
): Ref<T> {
  return useCallback((value: T) => {
    for (const ref of refs) {
      if (typeof ref === 'function') {
        ref(value);
      } else if (ref != null) {
        (ref as MutableRefObject<T>).current = value;
      }
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, refs);
}
