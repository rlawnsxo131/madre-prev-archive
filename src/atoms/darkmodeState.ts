import { atom, selector } from 'recoil';

interface DarkmodeState {
  them: 'dark' | 'light';
}

const darkmodeState = atom<DarkmodeState>({
  key: 'darkmodeState',
  default: {
    them: 'light',
  },
});

export const darkmodeSelector = selector<DarkmodeState>({
  key: 'darkmodeSelector',
  get: ({ get }) => get(darkmodeState),
  set: ({ set }, newValue) => set(darkmodeState, (prevValue) => newValue),
});
