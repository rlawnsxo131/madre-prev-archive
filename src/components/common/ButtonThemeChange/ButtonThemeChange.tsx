import { css } from '@emotion/react';
import useThemeActions from '../../../hooks/theme/useThemeActions';
import useThemeState from '../../../hooks/theme/useThemeState';
import { LightIcon, NightIcon } from '../../../image/icons';
import { Theme } from '../../../store/theme';
import { themeColor } from '../../../styles';

interface ButtonThemeChangeProps {}

function ButtonThemeChange(props: ButtonThemeChangeProps) {
  const { theme } = useThemeState();
  const { handleTheme } = useThemeActions();

  return (
    <button css={block(theme)} onClick={handleTheme}>
      {theme === 'light' && <LightIcon />}
      {theme === 'dark' && <NightIcon />}
    </button>
  );
}

const block = (theme: Theme) => css`
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

export default ButtonThemeChange;
