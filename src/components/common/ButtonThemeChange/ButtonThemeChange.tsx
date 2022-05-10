import { css } from '@emotion/react';
import useThemeActions from '../../../hooks/theme/useThemeActions';
import useThemeState from '../../../hooks/theme/useThemeState';
import { LightIcon, NightIcon } from '../../../image/icons';
import { themePalette } from '../../../styles';

interface ButtonThemeChangeProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement> {}

function ButtonThemeChange(props: ButtonThemeChangeProps) {
  const { theme } = useThemeState();
  const { handleTheme } = useThemeActions();

  return (
    <button css={block} onClick={handleTheme}>
      {theme === 'light' && <LightIcon />}
      {theme === 'dark' && <NightIcon />}
    </button>
  );
}

const block = css`
  background: inherit;
  border: none;
  box-shadow: none;
  border-radius: 0;
  padding: 0;
  overflow: visible;
  display: flex;
  align-items: center;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  svg {
    fill: ${themePalette.fill1};
  }
`;

export default ButtonThemeChange;
