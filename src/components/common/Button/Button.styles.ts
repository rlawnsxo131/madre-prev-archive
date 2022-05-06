import { palette } from '../../../styles';
import { themePalette } from '../../../styles';

export type ButtonShape = 'rect' | 'round';
export type ButtonSize = 'small' | 'medium' | 'large' | 'responsive';
export type ButtonColor = 'gray' | 'blue';

export const buttonSizeMap = {
  small: {
    height: '1.5rem',
    padding: '0 0.9375rem',
    fontSize: '0.75rem',
  },
  medium: {
    height: '2rem',
    padding: '0 1.25rem',
    fontSize: '1rem',
  },
  large: {
    height: '2.5rem',
    padding: '0 1.125rem',
    fontSize: '1.125rem',
  },
};

export const buttonColorMap = {
  gray: {
    font: palette.white,
    background: palette.gray['700'],
    hoverBackground: palette.gray['600'],
  },
  blue: {
    font: palette.white,
    background: palette.blue['600'],
    hoverBackground: palette.blue['500'],
  },
};

export const buttonOutlineStyle = {
  blue: {
    default: palette.blue['600'],
    hover: palette.blue['500'],
  },
  gray: {
    default: palette.gray['700'],
    hover: palette.gray['600'],
  },
  background: 'inherit',
  opacity: themePalette.opacity_button1,
};
