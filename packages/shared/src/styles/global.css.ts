import { createGlobalTheme, globalStyle } from '@vanilla-extract/css';

import { darkColors, lightColors, themes } from './themes';

createGlobalTheme(`:root, :root[data-theme="light"]`, themes, {
  color: { ...lightColors },
});
createGlobalTheme(`:root[data-theme="dark"]`, themes, {
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

globalStyle('*', {
  boxSizing: 'inherit',
  fontFamily: 'inherit',
});

globalStyle('a', {
  textDecoration: 'none',
});

globalStyle(`:root, :root[data-theme="dark"] :where(body, #root)`, {
  backgroundColor: themes.color['white'],
});

globalStyle(`:root[data-theme="dark"] :where(body, #root)`, {
  backgroundColor: themes.color['gray2'],
});

globalStyle('h1, h2, h3, h4, h5, h6, p', {
  color: themes.color['gray12'],
});
