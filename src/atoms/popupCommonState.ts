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
  const setState = useSetRecoilState(popupCommonStateSelector);
  const resetState = useResetRecoilState(popupCommonStateSelector);

  const showPopup = useCallback(
    ({ title, message }: { title: string; message: string }) => {
      setState((prev) => ({
        ...prev,
        title,
        message,
        visible: true,
      }));
    },
    [setState],
  );

  const closePopup = useCallback(() => {
    resetState();
  }, [resetState]);

  return {
    state,
    showPopup,
    closePopup,
  };
}
