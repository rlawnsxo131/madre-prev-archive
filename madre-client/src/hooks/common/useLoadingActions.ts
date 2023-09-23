import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import common from '../../store/common';

export default function useLoadingActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      show() {
        dispatch(common.actions.showLoading());
      },
      close() {
        dispatch(common.actions.closeLoading());
      },
    }),
    [dispatch],
  );
}
