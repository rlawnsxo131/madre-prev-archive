import { css } from '@emotion/react';
import { useSelector } from 'react-redux';
import { LightIcon, NightIcon } from '../../../image/icons';
import { RootState } from '../../../store';
import { Theme } from '../../../store/theme';
import { themeColor } from '../../../styles';
import useSetTheme from '../../../hooks/theme/useSetTheme';

interface ThemeButtonProps {}

function ThemeButton(props: ThemeButtonProps) {
  const theme = useSelector((state: RootState) => state.theme.theme);
  const onClick = useSetTheme();

  return (
    <button css={block(theme)} onClick={onClick}>
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
