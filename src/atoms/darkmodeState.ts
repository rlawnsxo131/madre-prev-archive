import { useCallback } from 'react';
import { atom, useRecoilValue, useSetRecoilState } from 'recoil';
import { MADRE_DARKMODE } from '../constants';
import { storage } from '../lib/storage';

export type DarkmodeTheme = 'dark' | 'light';
interface DarkmodeState {
  theme: DarkmodeTheme;
}

export const darkmodeState = atom<DarkmodeState>({
  key: 'darkmodeState',
  default: {
    theme: 'light',
  },
});

export function useDarkmodeValue() {
  return useRecoilValue(darkmodeState);
}

export function useDarkmodeActions() {
  const set = useSetRecoilState(darkmodeState);

  const handleDarkmode = useCallback(() => {
    set((prev) => {
      const theme = prev.theme === 'dark' ? 'light' : 'dark';
      storage.setItem(MADRE_DARKMODE, theme);
      return {
        ...prev,
        theme,
      };
    });
  }, [set]);

  return {
    handleDarkmode,
  };
}
