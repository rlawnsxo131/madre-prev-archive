import { cyan, cyanDark, gray, grayDark, red, redDark } from '@radix-ui/colors';
import { createGlobalThemeContract } from '@vanilla-extract/css';

const baseColors = {
  black: '#000000',
  white: '#FFFFFF',
};

export const lightColors = Object.assign({ ...baseColors }, cyan, gray, red);
export const darkColors = Object.assign(
  { ...baseColors },
  cyanDark,
  grayDark,
  redDark,
);

type ColorKeys = keyof typeof lightColors;
type Keys = Record<ColorKeys, ColorKeys>;

const keys = (Object.keys(lightColors) as ColorKeys[]).reduce<Keys>(
  (acc, key) => {
    acc[key] = key;
    return acc;
  },
  {} as Keys,
);

export const vars = createGlobalThemeContract(
  {
    color: {
      ...keys,
    },
  },
  (value) => `--color-${value}`,
);
