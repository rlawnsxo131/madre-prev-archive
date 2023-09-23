import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import common from '../../store/common';

interface PopupCommonshowParams {
  title?: string;
  message?: string;
}

export default function usePopupCommonActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      show({ title = '', message = '' }: PopupCommonshowParams) {
        dispatch(
          common.actions.showPopupCommon({
            title,
            message,
          }),
        );
      },
      close() {
        dispatch(common.actions.closePopupCommon());
      },
    }),
    [dispatch],
  );
}
