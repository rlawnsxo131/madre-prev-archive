import { useCallback } from 'react';
import {
  atom,
  DefaultValue,
  selector,
  useRecoilValue,
  useSetRecoilState,
} from 'recoil';

export type DarkmodeThemeType = 'dark' | 'light';
interface DarkmodeState {
  theme: DarkmodeThemeType;
}

const darkmodeState = atom<DarkmodeState>({
  key: 'darkmodeState',
  default: {
    theme: 'light',
  },
});

export const darkmodeStateSelector = selector<DarkmodeState>({
  key: 'darkmodeStateSelector',
  get: ({ get }) => get(darkmodeState),
  set: ({ set }, newValue) =>
    set(darkmodeState, (prevValue) =>
      newValue instanceof DefaultValue
        ? newValue
        : { ...prevValue, ...newValue },
    ),
});

export function useDarkmode() {
  const state = useRecoilValue(darkmodeStateSelector);
  const set = useSetRecoilState(darkmodeStateSelector);

  const handleDarkmode = useCallback(() => {
    set((prev) => ({
      ...prev,
      theme: prev.theme === 'dark' ? 'light' : 'dark',
    }));
  }, [set]);

  return {
    state,
    handleDarkmode,
  };
}
