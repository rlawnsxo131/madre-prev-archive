import { MutableSnapshot } from 'recoil';
import { MADRE_COLOR_THEME } from '../constants';
import { storage } from '../lib/storage';
import { getPrefersColorScheme } from '../lib/utils';
import { colorThemeState } from './colorThemeState';

export default function recoilInitializer({ set }: MutableSnapshot) {
  const theme = storage.getItem(MADRE_COLOR_THEME) ?? getPrefersColorScheme();
  set(colorThemeState, { theme });
}
