interface ThemeVariables {
  bg1: string;
  bg2: string;
  bg3: string;
  bg4: string;

  bg_element1: string;

  bg_button1: string;

  outline_button_gray: string;
  outline_button_gray_hover: string;
  outline_button_disabled: string;

  text1: string;
  text2: string;

  opacity_button1: string;

  border_element1: string;

  fill1: string;

  shadow1: string;

  opaque_layer: string;
}

type Theme = 'light' | 'dark';
type VariableKey = keyof ThemeVariables;
type ThemedPalette = Record<VariableKey, string>;

/**
 * border_element1: Input, ScreenSignUp
 * bg_element1: PopupBase, PopupAuth, HeaderMobileNavigationLinks
 * shadow1: HeaderMobileNavigationLinks
 */
const themeVariableSets: Record<Theme, ThemeVariables> = {
  light: {
    bg1: '#FFFF',
    bg2: '#f8f9fa',
    bg3: '#FFFF',
    bg4: '#FFFF',

    bg_element1: '#FFFF',

    bg_button1: 'rgba(255, 255, 255, 0)',

    outline_button_gray: '#495057',
    outline_button_gray_hover: '#868e96',
    outline_button_disabled: '#adb5bd',

    text1: '#242526',
    text2: '#495057',

    opacity_button1: '0.7',

    border_element1: '#e9ecef',

    fill1: '#242526',

    shadow1: '1px 1px 3px 1px #adb5bd',

    opaque_layer: '',
  },
  dark: {
    bg1: '#B0B3B8',
    bg2: '#3A3B3C',
    bg3: '#242526',
    bg4: '#18191A',

    bg_element1: '#242526',

    bg_button1: '#FFFF',

    outline_button_gray: '#adb5bd',
    outline_button_gray_hover: '#ced4da',
    outline_button_disabled: '#495057',

    text1: '#FFFF',
    text2: '#adb5bd',

    opacity_button1: '1',

    border_element1: '#868e96',

    fill1: '#FFFF',

    shadow1: '1px 1px 3px 1px #18191A',

    opaque_layer: '',
  },
};

const buildCssVariables = (variables: ThemeVariables) => {
  const keys = Object.keys(variables) as (keyof ThemeVariables)[];
  return keys.reduce(
    (acc, key) =>
      acc.concat(`--${key.replace(/_/g, '-')}: ${variables[key]};`, '\n'),
    '',
  );
};

export const themes = {
  light: buildCssVariables(themeVariableSets.light),
  dark: buildCssVariables(themeVariableSets.dark),
};

const cssVar = (name: string) => `var(--${name.replace(/_/g, '-')})`;

const variableKeys = Object.keys(themeVariableSets.light) as VariableKey[];

export const themePalette: Record<VariableKey, string> = variableKeys.reduce(
  (acc, current) => {
    acc[current] = cssVar(current);
    return acc;
  },
  {} as ThemedPalette,
);
