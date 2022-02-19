import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../redux';
import { setPopupLogin } from '../../redux/core';

export default function usePopupLoginActions() {
  const dispatch = useDispatch<AppDispatch>();
  return useMemo(
    () => ({
      onShow() {
        dispatch(
          setPopupLogin({
            visible: true,
          }),
        );
      },
      onClose() {
        dispatch(
          setPopupLogin({
            visible: false,
          }),
        );
      },
    }),
    [],
  );
}
