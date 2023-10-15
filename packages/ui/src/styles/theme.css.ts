import { cyan, cyanDark, gray, grayDark, red, redDark } from '@radix-ui/colors';
import {
  createGlobalTheme,
  createGlobalThemeContract,
} from '@vanilla-extract/css';

const lightColors = Object.assign({}, cyan, gray, red);
const darkColors = Object.assign({}, cyanDark, grayDark, redDark);

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

createGlobalTheme('[data-theme="light"]', { ...lightColors });
createGlobalTheme('[data-theme="dark"]', { ...darkColors });
