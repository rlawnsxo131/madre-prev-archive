import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import popupAuth from '../../store/popupAuth';

export default function usePopupAuthActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      show() {
        dispatch(
          popupAuth.actions.setVisible({
            visible: true,
          }),
        );
      },
      close() {
        dispatch(
          popupAuth.actions.setVisible({
            visible: false,
          }),
        );
      },
      setError() {
        dispatch(
          popupAuth.actions.setIsError({
            isError: true,
          }),
        );
      },
      resetError() {
        dispatch(
          popupAuth.actions.setIsError({
            isError: false,
          }),
        );
      },
    }),
    [dispatch],
  );
}
