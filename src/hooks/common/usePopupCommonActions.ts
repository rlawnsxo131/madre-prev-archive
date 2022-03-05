import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import { setPopupCommon } from '../../store/common';

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
          setPopupCommon({
            visible: true,
            title,
            message,
          }),
        );
      },
      close() {
        dispatch(
          setPopupCommon({
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
