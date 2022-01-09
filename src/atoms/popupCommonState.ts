import { useCallback } from 'react';
import {
  atom,
  useRecoilValue,
  useResetRecoilState,
  useSetRecoilState,
} from 'recoil';

interface PopupCommonState {
  title: string;
  message: string;
  visible: boolean;
}

const popupCommonState = atom<PopupCommonState>({
  key: 'popupCommonState',
  default: {
    title: '',
    message: '',
    visible: false,
  },
});

export function usePopupCommonValue() {
  return useRecoilValue(popupCommonState);
}

export function usePopupCommonActions() {
  const set = useSetRecoilState(popupCommonState);
  const reset = useResetRecoilState(popupCommonState);

  const showPopup = useCallback(
    ({ title = '', message }: { title?: string; message: string }) => {
      set((prev) => ({
        ...prev,
        title,
        message,
        visible: true,
      }));
    },
    [set],
  );

  const closePopup = useCallback(() => {
    reset();
  }, [reset]);

  return {
    showPopup,
    closePopup,
  };
}
