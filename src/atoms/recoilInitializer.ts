import { MutableSnapshot } from 'recoil';
import { MADRE_DARKMODE } from '../constants';
import { storage } from '../lib/storage';
import { getPrefersColorScheme } from '../lib/utils';
import { darkmodeState } from './darkmodeState';

export default function recoilInitializer({ set }: MutableSnapshot) {
  const theme = storage.getItem(MADRE_DARKMODE) ?? getPrefersColorScheme();
  set(darkmodeState, { theme });
}
