import { atom, DefaultValue, selector, useRecoilValue } from 'recoil';

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

const toastStateSelector = selector<ToastState>({
  key: 'toastStateSelector',
  get: ({ get }) => get(toastState),
  set: ({ set }, newValue) =>
    set(toastState, (prevValue) =>
      newValue instanceof DefaultValue
        ? newValue
        : { ...prevValue, ...newValue },
    ),
});

export function useToastState() {
  return useRecoilValue(toastStateSelector);
}
