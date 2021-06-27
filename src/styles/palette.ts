/**
 * undraw color: violet['7']
 */
const palette = {
  white: '#FFFF',
  black: '#000000',
  gray: {
    0: '#f8f9fa',
    1: '#f1f3f5',
    2: '#e9ecef',
    3: '#dee2e6',
    4: '#ced4da',
    5: '#adb5bd',
    6: '#868e96',
    7: '#495057',
    8: '#343a40',
    9: '#212529',
  },
  red: {
    0: '#fff5f5',
    1: '#ffe3e3',
    2: '#ffc9c9',
    3: '#ffa8a8',
    4: '#ff8787',
    5: '#ff6b6b',
    6: '#fa5252',
    7: '#f03e3e',
    8: '#e03131',
    9: '#c92a2a',
  },
  violet: {
    0: '#f3f0ff',
    1: '#e5dbff',
    2: '#d0bfff',
    3: '#b197fc',
    4: '#9775fa',
    5: '#845ef7',
    6: '#7950f2',
    7: '#7048e8',
    8: '#6741d9',
    9: '#5f3dc4',
  },
};

const darkmodeColor = {
  background: {
    1: '#282c35',
    2: '#292f35',
  },
};

export const themeColor = {
  common: {
    light: palette.white,
    dark: darkmodeColor.background['1'],
  },
  commonDeepDark: {
    light: palette.white,
    dark: darkmodeColor.background['2'],
  },
  font: {
    light: palette.black,
    dark: palette.gray['3'],
  },
  fill: {
    light: palette.black,
    dark: palette.gray['3'],
  },
  hover: {
    light: palette.gray['3'],
    dark: darkmodeColor.background['2'],
  },
  shadow: {
    light: `1px 1px 3px 1px ${palette.gray['3']}`,
    dark: `1px 1px 3px 1px ${palette.gray['9']}`,
  },
};

export default palette;
