import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../redux';
import { setPopupAuth } from '../../redux/core';

export default function usePopupAuthActions() {
  const dispatch = useDispatch<AppDispatch>();
  return useMemo(
    () => ({
      onShow() {
        dispatch(
          setPopupAuth({
            visible: true,
          }),
        );
      },
      onClose() {
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
