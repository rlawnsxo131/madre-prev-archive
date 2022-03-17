import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import common from '../../store/common';

interface PopupCommshowParams {
  title?: string;
  message?: string;
}

export default function usePopupCommonActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      show({ title = '', message = '' }: PopupCommshowParams) {
        dispatch(
          common.actions.setPopupCommon({
            visible: true,
            title,
            message,
          }),
        );
      },
      close() {
        dispatch(
          common.actions.setPopupCommon({
            visible: false,
            title: '',
            message: '',
          }),
        );
      },
    }),
    [dispatch],
  );
}
