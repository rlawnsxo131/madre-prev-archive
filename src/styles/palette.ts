const palette = {
  white: '#FFFF',
  black: '#000000',
  gray: {
    50: '#f8f9fa',
    100: '#f1f3f5',
    200: '#e9ecef',
    300: '#dee2e6',
    400: '#ced4da',
    500: '#adb5bd',
    600: '#868e96',
    700: '#495057',
    800: '#343a40',
    900: '#212529',
  },
  red: {
    50: '#fff5f5',
    100: '#ffe3e3',
    200: '#ffc9c9',
    300: '#ffa8a8',
    400: '#ff8787',
    500: '#ff6b6b',
    600: '#fa5252',
    700: '#f03e3e',
    800: '#e03131',
    900: '#c92a2a',
  },
};

const darkmodeColor = {
  background: {
    100: '#282c35',
    200: '#292f35',
  },
};

export const themeColor = {
  background: {
    light: palette.white,
    dark: darkmodeColor.background['100'],
  },
  font: {
    light: palette.black,
    dark: palette.gray['200'],
  },
  fill: {
    light: palette.black,
    dark: palette.gray['200'],
  },
  hover: {
    light: palette.gray['200'],
    dark: darkmodeColor.background['200'],
  },
  shadow: {
    light: `1px 1px 3px 1px ${palette.gray['500']}`,
    dark: `1px 1px 3px 1px ${palette.black}`,
  },
};

export default palette;
