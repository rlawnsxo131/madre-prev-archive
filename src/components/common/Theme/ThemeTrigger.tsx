import { css } from '@emotion/react';
import { DarkmodeThemeType, useDarkmode } from '../../../atoms/darkmodeState';
import { LightIcon, NightIcon } from '../../../image/icons';
import { themeColor } from '../../../styles/palette';

interface ThemTriggerProps {}

function ThemeTrigger(props: ThemTriggerProps) {
  const { state, handleDarkmode } = useDarkmode();

  return (
    <button css={block(state.theme)} onClick={handleDarkmode}>
      {state.theme === 'light' && <LightIcon />}
      {state.theme === 'dark' && <NightIcon />}
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
