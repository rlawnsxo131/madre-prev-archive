import { darkmodeColor, palette } from '.';

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
