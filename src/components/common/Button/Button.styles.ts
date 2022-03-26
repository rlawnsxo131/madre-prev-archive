import { palette } from '../../../styles';
import { themePalette } from '../../../styles';

export type ButtonShape = 'rect' | 'round';
export type ButtonSize = 'small' | 'medium' | 'large' | 'responsive';
export type ButtonColor = 'gray' | 'pink';

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
    background: palette.gray['700'],
    hoverBackground: palette.gray['600'],
  },
  pink: {
    background: palette.pink['600'],
    hoverBackground: palette.pink['500'],
  },
};

export const outlineButtonStyle = {
  background: themePalette.bg_button1,
  opacity: themePalette.opacity_button1,
};
