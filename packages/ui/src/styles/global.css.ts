import { createGlobalTheme, globalStyle } from '@vanilla-extract/css';

import { dataThemeSelector } from './selector';
import { darkColors, lightColors, vars, zIndices } from './vars.css';

createGlobalTheme(`:root, :root${dataThemeSelector.light}`, vars, {
  color: { ...lightColors },
  zIndices: { ...zIndices },
});
createGlobalTheme(`:root${dataThemeSelector.dark}`, vars, {
  color: { ...darkColors },
  zIndices: { ...zIndices },
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
  background: vars.color.white,
});

globalStyle(`:root${dataThemeSelector.dark} :where(body, #root)`, {
  background: vars.color.gray2,
});

globalStyle('h1, h2, h3, h4, h5, h6, p', {
  color: vars.color.gray12,
});
