import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { MADRE_COLOR_THEME } from '../constants';
import { Storage } from '../lib/storage';

export type Theme = 'light' | 'dark';

interface ThemeState {
  theme: Theme;
}

const initialState: ThemeState = {
  theme: 'light',
};

const theme = createSlice({
  name: 'theme',
  initialState,
  reducers: {
    handleTheme(state) {
      const currentTheme = state.theme === 'light' ? 'dark' : 'light';
      Storage.setItem(MADRE_COLOR_THEME, currentTheme);
      state.theme = currentTheme;
    },
    setTheme(state, action: PayloadAction<Theme>) {
      Storage.setItem(MADRE_COLOR_THEME, action.payload);
      state.theme = action.payload;
    },
  },
});

export default theme;
