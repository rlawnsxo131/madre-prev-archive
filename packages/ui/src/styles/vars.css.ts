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

export const zIndices = {
  overlay: '12',
};

function generateKeyObject<Obj extends Record<string, string>>(
  obj: Obj,
): Record<keyof Obj, keyof Obj> {
  return Object.keys(obj).reduce(
    (acc, key) => {
      acc[key as keyof Obj] = key as keyof Obj;
      return acc;
    },
    {} as Record<keyof Obj, keyof Obj>,
  );
}

/**
 * @TODO
 * env(safe-area-inset-top)
 * env(safe-area-inset-right)
 * env(safe-area-inset-bottom)
 * env(safe-area-inset-left)
 */
export const vars = createGlobalThemeContract(
  {
    color: {
      ...generateKeyObject(lightColors),
    },
    zIndices: {
      ...generateKeyObject(zIndices),
    },
  },
  (_, path) => {
    const [key, val] = path;
    return `--${key}-${val}`;
  },
);
