import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import popupAuth from '../../store/popupAuth';

export default function usePopupAuthActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      show() {
        dispatch(popupAuth.actions.show());
      },
      close() {
        dispatch(popupAuth.actions.close());
      },
      setError() {
        dispatch(popupAuth.actions.setIsError());
      },
      resetError() {
        dispatch(popupAuth.actions.resetIsError());
      },
    }),
    [dispatch],
  );
}
