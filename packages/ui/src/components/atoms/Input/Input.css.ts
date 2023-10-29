import { style } from '@vanilla-extract/css';

import { themes } from '@/styles';

/**
 * base
 */
export const block = style({
  background: 'transparent',
  outline: 'none',
  border: `1px solid ${themes.color['gray8']}`,
  borderRadius: '0.3rem',
  caretColor: themes.color['cyan10'],
  color: themes.color['gray12'],
  margin: '1rem',
  selectors: {
    '&:focus': {
      border: `1px solid ${themes.color['cyan10']}`,
    },
  },
});

/**
 * size
 */

/**
 * theme
 */
