import { style, styleVariants } from '@vanilla-extract/css';
import { themes } from 'src/styles';

/**
 * base
 */
export const base = style({
  background: 'transparent',
  outline: 'none',
  borderRadius: '0.25rem',
  border: `1px solid ${themes.color['gray8']}`,
  caretColor: themes.color['cyan10'],
  color: themes.color['gray12'],
  selectors: {
    '&:focus': {
      border: `2px solid ${themes.color['cyan10']}`,
    },
    '&:-webkit-autofill, &:-webkit-autofill:hover, &:-webkit-autofill:focus, &:-webkit-autofill:active':
      {
        transition: 'background-color 5000s ease-in-out 0s',
        WebkitTransition: 'background-color 9999s ease-out',
        WebkitBoxShadow: '0 0 0px 1000px transparent inset !important',
        WebkitTextFillColor: 'transparent !important',
      },
  },
});

/**
 * size
 */
export const size = styleVariants({
  small: {
    height: '1.725rem',
    lineHeight: '1.725rem',
    padding: '0.25rem 0.5rem',
    fontSize: '0.825rem',
  },
  medium: {
    height: '2rem',
    lineHeight: '2rem',
    padding: '0.25rem 0.5rem',
    fontSize: '1rem',
  },
  large: {
    height: '2.275rem',
    lineHeight: '2.275rem',
    padding: '0.25rem 0.5rem',
    fontSize: '1.25rem',
  },
});

/**
 * warn
 */
export const warn = style({
  border: `1px solid ${themes.color['red9']}`,
  caretColor: themes.color['red9'],
  color: themes.color['red9'],
  selectors: {
    '&:focus': {
      border: `2px solid ${themes.color['red9']}`,
    },
  },
});
