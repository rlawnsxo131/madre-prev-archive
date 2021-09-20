import { css } from '@emotion/react';
import {
  DarkmodeThemeType,
  useDarkmodeValue,
  useDarkmodeActions,
} from '../../../atoms/darkmodeState';
import { LightIcon, NightIcon } from '../../../image/icons';
import { themeColor } from '../../../styles/palette';

interface ThemTriggerProps {}

function ThemeTrigger(props: ThemTriggerProps) {
  const { theme } = useDarkmodeValue();
  const { handleDarkmode } = useDarkmodeActions();

  return (
    <button css={block(theme)} onClick={handleDarkmode}>
      {theme === 'light' && <LightIcon />}
      {theme === 'dark' && <NightIcon />}
    </button>
  );
}

const block = (theme: DarkmodeThemeType) => css`
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

export default ThemeTrigger;
