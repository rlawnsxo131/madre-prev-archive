import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import { setLoading } from '../../store/core';

export default function useLoadingActions() {
  const dispatch = useDispatch<AppDispatch>();
  return useMemo(
    () => ({
      show() {
        dispatch(
          setLoading({
            visible: true,
          }),
        );
      },
      close() {
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
