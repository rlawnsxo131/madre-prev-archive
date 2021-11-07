import { MutableSnapshot } from 'recoil';
import { MADRE_COLOR_THEME } from '../constants';
import { Storage } from '../lib/storage';
import { getPrefersColorScheme } from '../lib/utils';
import { colorThemeState } from './colorThemeState';

export default function recoilInitializer({ set }: MutableSnapshot) {
  const theme = Storage.getItem(MADRE_COLOR_THEME) ?? getPrefersColorScheme();
  set(colorThemeState, { theme });
}
