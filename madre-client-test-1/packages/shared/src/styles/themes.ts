import { cyan, cyanDark, gray, grayDark, red, redDark } from '@radix-ui/colors';
import { createGlobalThemeContract } from '@vanilla-extract/css';

const baseColors = {
  black: '#000000',
  white: '#FFFFFF',
  overlay: 'rgba(100, 100, 100, 0.5)',
};
const customColors = {
  light: {
    'bg-element1': '#FFFF',
  },
  dark: {
    'bg-element1': '#242526',
  },
};

export const lightColors = Object.assign(
  {
    ...baseColors,
    ...customColors.light,
  },
  cyan,
  gray,
  red,
);
export const darkColors = Object.assign(
  {
    ...baseColors,
    ...customColors.dark,
  },
  cyanDark,
  grayDark,
  redDark,
);

type ColorKey = keyof typeof lightColors;

export const themes = createGlobalThemeContract(
  {
    color: {
      ...(Object.keys(lightColors) as ColorKey[]).reduce(
        (acc, key) => {
          acc[key] = key;
          return acc;
        },
        {} as Record<ColorKey, ColorKey>,
      ),
    },
  },
  (_, path) => {
    const [key, val] = path;
    return `--${key}-${val}`;
  },
);
