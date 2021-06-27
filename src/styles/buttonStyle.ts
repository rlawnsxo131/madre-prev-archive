import palette from './palette';

export type ButtonShapeType = 'rect' | 'round';
export type ButtonSizeType = 'small' | 'medium' | 'large' | 'responsive';
export type ButtonColorType = 'gray' | 'red' | 'violet';

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
    background: palette.gray[7],
    hoverBackground: palette.gray[6],
  },
  red: {
    background: palette.red[7],
    hoverBackground: palette.red[5],
  },
};

export const outlineButtonStyle = {
  light: {
    background: 'rgba(255, 255, 255, 0)',
    opacity: '0.7',
  },
  dark: {
    background: palette.white,
    opacity: '1',
  },
};
