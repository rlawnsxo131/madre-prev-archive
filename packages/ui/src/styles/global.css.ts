import { createGlobalTheme, globalStyle } from '@vanilla-extract/css';

import { dataThemeSelector } from './selector';
import { darkColors, lightColors, themes } from './themes';

createGlobalTheme(`:root, :root${dataThemeSelector.light}`, themes, {
  color: { ...lightColors },
});
createGlobalTheme(`:root${dataThemeSelector.dark}`, themes, {
  color: { ...darkColors },
});

globalStyle('html, body, #root', {
  margin: 0,
  padding: 0,
  height: '100%',
  boxSizing: 'border-box',
  fontFamily: 'Inter, system-ui, Avenir, Helvetica, Arial, sans-serif',
  colorScheme: 'light dark',
  fontSynthesis: 'none',
  textRendering: 'optimizeLegibility',
  WebkitFontSmoothing: 'antialiased',
  MozOsxFontSmoothing: 'grayscale',
  WebkitTextSizeAdjust: '100%',
});

globalStyle('html, body, #root *', {
  boxSizing: 'inherit',
});

globalStyle(`:root, :root${dataThemeSelector.light} :where(body, #root)`, {
  backgroundColor: themes.color['white'],
});

globalStyle(`:root${dataThemeSelector.dark} :where(body, #root)`, {
  backgroundColor: themes.color['gray2'],
});

globalStyle('h1, h2, h3, h4, h5, h6, p', {
  color: themes.color['gray12'],
});

globalStyle('a', {
  all: 'unset',
});
