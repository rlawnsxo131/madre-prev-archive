import { style, styleVariants } from '@vanilla-extract/css';

import { themes } from '@/styles';

/**
 * base
 */

export const block = style({
  display: 'inline-flex',
  justifyContent: 'center',
  alignItems: 'center',
  cursor: 'pointer',
  transition: 'all .15s',
  selectors: {
    '& + &': {
      marginLeft: '1rem',
    },
    '&:active:not([disabled])': {
      transform: 'scale(.99, .98)',
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

/**
 * solid
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
    selectors: {
      '&:hover:not([disabled])': {
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
    selectors: {
      '&:hover:not([disabled])': {
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
    selectors: {
      '&:hover:not([disabled])': {
        backgroundColor: themes.color['red11'],
        border: `1px solid ${themes.color['red11']}`,
      },
      '&:disabled': {
        ...solidDisabled,
      },
    },
  },
});

/**
 * outline
 */
const outlineDisabled = {
  color: themes.color['gray6'],
  border: `1px solid ${themes.color['gray6']}`,
};

const themeOutline = styleVariants({
  primary: {
    color: themes.color['cyan11'],
    backgroundColor: 'transparent',
    border: `1px solid ${themes.color['cyan10']}`,
    selectors: {
      '&:hover:not([disabled])': {
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
    selectors: {
      '&:hover:not([disabled])': {
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
    selectors: {
      '&:hover:not([disabled])': {
        backgroundColor: themes.color['red3'],
      },
      '&:disabled': {
        ...outlineDisabled,
      },
    },
  },
});

/**
 * ghost
 */
const ghostDisabled = {
  color: themes.color['gray6'],
};

const themeGhost = styleVariants({
  primary: {
    color: themes.color['cyan11'],
    backgroundColor: 'transparent',
    border: 'transparent',
    selectors: {
      '&:hover:not([disabled])': {
        backgroundColor: themes.color['cyan3'],
      },
      '&:disabled': {
        ...ghostDisabled,
      },
    },
  },
  secondary: {
    color: themes.color['cyan11'],
    backgroundColor: 'transparent',
    border: 'transparent',
    selectors: {
      '&:hover:not([disabled])': {
        backgroundColor: themes.color['gray3'],
      },
      '&:disabled': {
        ...ghostDisabled,
      },
    },
  },
  warn: {
    color: themes.color['red11'],
    backgroundColor: 'transparent',
    border: 'transparent',
    selectors: {
      '&:hover:not([disabled])': {
        backgroundColor: themes.color['red3'],
      },
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
