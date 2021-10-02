import { useCallback } from 'react';
import { atom, useRecoilValue, useSetRecoilState } from 'recoil';
import { MADRE_COLOR_THEME } from '../constants';
import { storage } from '../lib/storage';

export type ColorTheme = 'dark' | 'light';
interface ColorThemeState {
  theme: ColorTheme;
}

export const colorThemeState = atom<ColorThemeState>({
  key: 'colorThemeState',
  default: {
    theme: 'light',
  },
});

export function useColorThemeValue() {
  return useRecoilValue(colorThemeState);
}

export function useColorThemeActions() {
  const set = useSetRecoilState(colorThemeState);

  const handleColorTheme = useCallback(() => {
    set((prev) => {
      const theme = prev.theme === 'dark' ? 'light' : 'dark';
      storage.setItem(MADRE_COLOR_THEME, theme);
      return {
        ...prev,
        theme,
      };
    });
  }, [set]);

  return {
    handleColorTheme,
  };
}
