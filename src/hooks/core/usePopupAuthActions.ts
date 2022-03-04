import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import { setPopupAuth } from '../../store/core';

export default function usePopupAuthActions() {
  const dispatch = useDispatch<AppDispatch>();
  return useMemo(
    () => ({
      show() {
        dispatch(
          setPopupAuth({
            visible: true,
          }),
        );
      },
      close() {
        dispatch(
          setPopupAuth({
            visible: false,
          }),
        );
      },
    }),
    [dispatch],
  );
}
