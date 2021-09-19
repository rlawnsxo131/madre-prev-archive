import { atom, useRecoilValue } from 'recoil';

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

export function useToastState() {
  return useRecoilValue(toastState);
}
