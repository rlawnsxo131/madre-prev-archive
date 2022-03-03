import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import { setLoading } from '../../store/core';

export default function useLoadingActions() {
  const dispatch = useDispatch<AppDispatch>();
  return useMemo(
    () => ({
      onShow() {
        dispatch(
          setLoading({
            visible: true,
          }),
        );
      },
      onClose() {
        dispatch(
          setLoading({
            visible: false,
          }),
        );
      },
    }),
    [dispatch],
  );
}
