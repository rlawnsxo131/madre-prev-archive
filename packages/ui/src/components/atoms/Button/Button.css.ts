import { style, styleVariants } from '@vanilla-extract/css';

import { themes } from '@/styles';

export const block = style({
  display: 'inline-flex',
  justifyContent: 'center',
  alignItems: 'center',
  cursor: 'pointer',
});

export const size = styleVariants({
  small: {
    height: '1.725rem',
    fontSize: '0.825rem',
    padding: '0.25rem 0.725rem',
  },
  medium: {
    height: '2rem',
    fontSize: '1rem',
    padding: '0.25rem 0.725rem',
  },
  large: {
    height: '2.5rem',
    fontSize: '1.125rem',
    padding: '0.25rem 1.125rem',
  },
});

export const fullWidth = style({
  flex: '1',
  width: '100%',
});

export const radius = styleVariants({
  none: {
    borderRadius: '0',
  },
  medium: {
    borderRadius: '0.5rem',
  },
  full: {
    borderRadius: '2rem',
  },
});

const themeSolid = styleVariants({
  primary: {
    color: themes.color['white'],
    background: themes.color['cyan10'],
    border: `1px solid ${themes.color['cyan10']}`,
    outlineColor: themes.color['cyan10'],
    selectors: {
      '&:hover': {
        background: themes.color['cyan11'],
        border: `1px solid ${themes.color['cyan11']}`,
        outlineColor: themes.color['cyan11'],
      },
    },
  },
  secondary: {
    color: themes.color['white'],
    background: themes.color['gray10'],
    border: `1px solid ${themes.color['gray10']}`,
    outlineColor: themes.color['gray10'],
    selectors: {
      '&:hover': {
        background: themes.color['gray11'],
        border: `1px solid ${themes.color['gray11']}`,
        outlineColor: themes.color['gray11'],
      },
    },
  },
  warn: {
    color: themes.color['white'],
    background: themes.color['red9'],
    border: `1px solid ${themes.color['red9']}`,
    outlineColor: themes.color['red9'],
    selectors: {
      '&:hover': {
        background: themes.color['red11'],
        border: `1px solid ${themes.color['red11']}`,
        outlineColor: themes.color['red11'],
      },
    },
  },
});

const themeOutline = styleVariants({
  primary: {},
  secondary: {},
  warn: {},
});

export const theme = {
  solid: themeSolid,
  outline: themeOutline,
};
