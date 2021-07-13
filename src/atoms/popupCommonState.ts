import { useCallback } from 'react';
import {
  atom,
  DefaultValue,
  selector,
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

const popupCommonStateSelector = selector<PopupCommonState>({
  key: 'popupStateSelector',
  get: ({ get }) => get(popupCommonState),
  set: ({ set }, newValue) =>
    set(popupCommonState, (prevValue) =>
      newValue instanceof DefaultValue
        ? newValue
        : { ...prevValue, ...newValue },
    ),
});

export function usePopupCommon() {
  const state = useRecoilValue(popupCommonStateSelector);
  const set = useSetRecoilState(popupCommonStateSelector);
  const reset = useResetRecoilState(popupCommonStateSelector);

  const showPopup = useCallback(
    ({ title, message }: { title: string; message: string }) => {
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
    state,
    showPopup,
    closePopup,
  };
}
