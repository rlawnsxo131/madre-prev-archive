import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import {
  setPopupAuthIsError,
  setPopupAuthVisible,
} from '../../store/popupAuth';

export default function usePopupAuthActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      show() {
        dispatch(
          setPopupAuthVisible({
            visible: true,
          }),
        );
      },
      close() {
        dispatch(
          setPopupAuthVisible({
            visible: false,
          }),
        );
      },
      setError() {
        dispatch(
          setPopupAuthIsError({
            isError: true,
          }),
        );
      },
      resetError() {
        dispatch(
          setPopupAuthIsError({
            isError: false,
          }),
        );
      },
    }),
    [dispatch],
  );
}
