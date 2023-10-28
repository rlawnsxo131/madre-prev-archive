import { keyframes, style, styleVariants } from '@vanilla-extract/css';

import { themes } from '@/styles';

/**
 * base
 */
const scale = keyframes({
  '0%': { transform: 'scale(1)' },
  '50%': { transform: 'scale(0.98)' },
  '100%': { transform: 'scale(1)' },
});

export const block = style({
  display: 'inline-flex',
  justifyContent: 'center',
  alignItems: 'center',
  cursor: 'pointer',
  selectors: {
    '& + &': {
      marginLeft: '1rem',
    },
    '&:active:not([disabled])': {
      animation: `${scale} .15s`,
    },
    '&:disabled': {
      cursor: 'not-allowed',
    },
  },
});

/**
 * size
 */
export const size = styleVariants({
  small: {
    height: '1.825rem',
    fontSize: '0.9rem',
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

/**
 * shape
 */
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

/**
 * theme
 */
const solidDisabled = {
  color: themes.color['gray1'],
  backgroundColor: themes.color['gray5'],
  border: `1px solid ${themes.color['gray5']}`,
};

const themeSolid = styleVariants({
  primary: {
    color: themes.color['white'],
    backgroundColor: themes.color['cyan10'],
    border: `1px solid ${themes.color['cyan10']}`,
    outlineColor: themes.color['cyan10'],
    selectors: {
      '&:hover': {
        backgroundColor: themes.color['cyan11'],
        border: `1px solid ${themes.color['cyan11']}`,
      },
      '&:disabled': {
        ...solidDisabled,
      },
    },
  },
  secondary: {
    color: themes.color['cyan11'],
    backgroundColor: themes.color['cyan4'],
    border: `1px solid ${themes.color['cyan4']}`,
    outlineColor: themes.color['cyan4'],
    selectors: {
      '&:hover': {
        backgroundColor: themes.color['cyan5'],
        border: `1px solid ${themes.color['cyan5']}`,
      },
      '&:disabled': {
        ...solidDisabled,
      },
    },
  },
  warn: {
    color: themes.color['white'],
    backgroundColor: themes.color['red9'],
    border: `1px solid ${themes.color['red9']}`,
    outlineColor: themes.color['red9'],
    selectors: {
      '&:hover': {
        backgroundColor: themes.color['red11'],
        border: `1px solid ${themes.color['red11']}`,
      },
      '&:disabled': {
        ...solidDisabled,
      },
    },
  },
});

const outlineDisabled = {
  color: themes.color['gray6'],
  backgroundColor: 'transparent',
  border: `1px solid ${themes.color['gray6']}`,
};

const themeOutline = styleVariants({
  primary: {
    color: themes.color['cyan11'],
    backgroundColor: 'transparent',
    border: `1px solid ${themes.color['cyan10']}`,
    outlineColor: themes.color['cyan10'],
    selectors: {
      '&:hover': {
        backgroundColor: themes.color['cyan2'],
        border: `1px solid ${themes.color['cyan11']}`,
      },
      '&:disabled': {
        ...outlineDisabled,
      },
    },
  },
  secondary: {
    color: themes.color['cyan11'],
    backgroundColor: themes.color['cyan2'],
    border: `1px solid ${themes.color['cyan10']}`,
    outlineColor: themes.color['cyan10'],
    selectors: {
      '&:hover': {
        backgroundColor: themes.color['cyan3'],
      },
      '&:disabled': {
        ...outlineDisabled,
      },
    },
  },
  warn: {
    color: themes.color['red11'],
    backgroundColor: themes.color['red2'],
    border: `1px solid ${themes.color['red10']}`,
    outlineColor: themes.color['red10'],
    selectors: {
      '&:hover': {
        backgroundColor: themes.color['red3'],
      },
      '&:disabled': {
        ...outlineDisabled,
      },
    },
  },
});

const ghostDisabled = {};

const themeGhost = styleVariants({
  primary: {
    selectors: {
      '&:hover': {},
      '&:disabled': {
        ...ghostDisabled,
      },
    },
  },
  secondary: {
    selectors: {
      '&:hover': {},
      '&:disabled': {
        ...ghostDisabled,
      },
    },
  },
  warn: {
    selectors: {
      '&:hover': {},
      '&:disabled': {
        ...ghostDisabled,
      },
    },
  },
});

export const theme = {
  solid: themeSolid,
  outline: themeOutline,
  ghost: themeGhost,
};
