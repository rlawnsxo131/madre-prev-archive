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

export const darkmodeState = atom<DarkmodeState>({
  key: 'darkmodeState',
  default: {
    theme: 'light',
  },
});

export function useDarkmodeState() {
  return useRecoilValue(darkmodeState);
}

export function useDarkmodeAction() {
  const set = useSetRecoilState(darkmodeState);

  const handleDarkmode = useCallback(() => {
    set((prev) => ({
      ...prev,
      theme: prev.theme === 'dark' ? 'light' : 'dark',
    }));
  }, [set]);

  return {
    handleDarkmode,
  };
}
