import { css } from '@emotion/react';
import { useSelector } from 'react-redux';
import { LightIcon, NightIcon } from '../../../image/icons';
import useTheme from '../../../lib/hooks/useTheme';
import { RootState } from '../../../store';
import { Theme } from '../../../store/theme';
import { themeColor } from '../../../styles';

interface ThemeButtonProps {}

function ThemeButton(props: ThemeButtonProps) {
  const theme = useSelector((state: RootState) => state.theme.theme);
  const { handleColorTheme } = useTheme();

  return (
    <button css={block(theme)} onClick={handleColorTheme}>
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

export default ThemeButton;
