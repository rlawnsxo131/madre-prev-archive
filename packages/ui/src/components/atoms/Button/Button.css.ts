import { style, styleVariants } from '@vanilla-extract/css';

import { vars } from '@/styles';

export const block = style({
  display: 'inline-flex',
  justifyContent: 'center',
  alignItems: 'center',
  cursor: 'pointer',
});

/**
 * @TODO 컬러 수정
 */
export const theme = styleVariants({
  primary: {
    color: vars.color.white,
    background: vars.color.cyan10,
    border: `1px solid ${vars.color.cyan10}`,
    outlineColor: vars.color.cyan10,
    selectors: {
      '&:hover': {
        background: vars.color.cyan11,
        border: `1px solid ${vars.color.cyan11}`,
        outlineColor: vars.color.cyan11,
      },
    },
  },
  secondary: {},
  ghost: {},
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

  responsive: {
    flex: '1',
    width: '100%',
    height: '2.5rem',
    fontSize: '1.125rem',
    padding: '0.25rem 1.125rem',
  },
});

export const radius = styleVariants({
  none: {},
  medium: {
    borderRadius: '0.5rem',
  },
  full: {
    borderRadius: '2rem',
  },
});
