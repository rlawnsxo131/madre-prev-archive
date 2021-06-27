import { atom, DefaultValue, selector } from 'recoil';

interface ToastState {
  message: string;
  visible: boolean;
}

const toastState = atom<ToastState>({
  key: 'toastState',
  default: {
    message: '',
    visible: false,
  },
});

export const toastStateSelector = selector<ToastState>({
  key: 'toastStateSelector',
  get: ({ get }) => get(toastState),
  set: ({ set }, newValue) =>
    set(toastState, (prevValue) =>
      newValue instanceof DefaultValue
        ? newValue
        : { ...prevValue, ...newValue },
    ),
});
