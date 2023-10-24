import { cyan, cyanDark, gray, grayDark, red, redDark } from '@radix-ui/colors';
import { createGlobalThemeContract } from '@vanilla-extract/css';

const baseColors = {
  black: '#000000',
  white: '#FFFFFF',
  overlay: 'rgba(100, 100, 100, 0.5)',
};
export const lightColors = Object.assign({ ...baseColors }, cyan, gray, red);
export const darkColors = Object.assign(
  { ...baseColors },
  cyanDark,
  grayDark,
  redDark,
);

export const themes = createGlobalThemeContract(
  {
    color: {
      ...Object.keys(lightColors).reduce(
        (acc, key: string) => {
          acc[key] = key;
          return acc;
        },
        {} as Record<string, string>,
      ),
    },
  },
  (_, path) => {
    const [key, val] = path;
    return `--${key}-${val}`;
  },
);
