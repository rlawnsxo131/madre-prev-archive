import { css } from '@emotion/react';
import {
  ColorTheme,
  useColorThemeValue,
  useColorThemeActions,
} from '../../../atoms/colorThemeState';
import { LightIcon, NightIcon } from '../../../image/icons';
import { themeColor } from '../../../styles';

interface ThemeButtonProps {}

function ThemeButton(props: ThemeButtonProps) {
  const { theme } = useColorThemeValue();
  const { handleColorTheme } = useColorThemeActions();

  return (
    <button css={block(theme)} onClick={handleColorTheme}>
      {theme === 'light' && <LightIcon />}
      {theme === 'dark' && <NightIcon />}
    </button>
  );
}

const block = (theme: ColorTheme) => css`
  background: inherit;
  border: none;
  box-shadow: none;
  border-radius: 0;
  overflow: visible;
  display: flex;
  align-items: center;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  svg {
    fill: ${themeColor.fill[theme]};
  }
  padding: 0.5rem;
`;

export default ThemeButton;
