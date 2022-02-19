import { palette } from '.';

const darkmodeColor = {
  deep: {
    100: '#242526',
    200: '#18191A',
  },
  medium: {
    100: '#B0B3B8',
    200: '#3A3B3C',
  },
  light: {
    100: '#FFFF',
    200: '#E4E6EB',
  },
  other: {
    100: '#282c35',
    200: '#292f35',
  },
};

const themeColor = {
  body: {
    light: darkmodeColor.light['100'],
    dark: darkmodeColor.deep['200'],
  },
  header: {
    light: darkmodeColor.light['100'],
    dark: darkmodeColor.deep['100'],
  },
  content: {
    light: darkmodeColor.light['100'],
    dark: darkmodeColor.deep['100'],
  },
  popup: {
    light: darkmodeColor.light['100'],
    dark: darkmodeColor.deep['100'],
  },
  screen: {
    light: darkmodeColor.light['200'],
    dark: darkmodeColor.deep['200'],
  },
  navigation: {
    light: darkmodeColor.light['100'],
    dark: darkmodeColor.deep['100'],
  },
  font: {
    light: darkmodeColor.deep['100'],
    dark: darkmodeColor.light['100'],
  },
  fill: {
    light: darkmodeColor.deep['100'],
    dark: darkmodeColor.light['100'],
  },
  shadow: {
    light: `1px 1px 3px 1px ${palette.gray['500']}`,
    dark: `1px 1px 3px 1px ${darkmodeColor.deep['200']}`,
  },
};

export default themeColor;
