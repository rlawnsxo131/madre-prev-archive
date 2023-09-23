import { css } from '@emotion/react';
import useThemeActions from '../../../hooks/theme/useThemeActions';
import useThemeState from '../../../hooks/theme/useThemeState';
import { LightIcon, NightIcon } from '../../../image/icons';
import { basicStyles } from '../../../styles';

interface ButtonThemeChangeProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement> {}

function ButtonThemeChange(props: ButtonThemeChangeProps) {
  const { theme } = useThemeState();
  const { handleTheme } = useThemeActions();

  return (
    <button css={[basicStyles.button, block]} onClick={handleTheme}>
      {theme === 'light' && <LightIcon />}
      {theme === 'dark' && <NightIcon />}
    </button>
  );
}

const block = css`
  background: inherit;
  display: flex;
  align-items: center;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
`;

export default ButtonThemeChange;
