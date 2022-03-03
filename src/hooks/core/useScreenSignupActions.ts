import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import { setScreenSignup } from '../../store/core';

export default function useScreenSignupActions() {
  const dispatch = useDispatch<AppDispatch>();
  return useMemo(
    () => ({
      onShow() {
        dispatch(setScreenSignup({ visible: true }));
      },
      onClose() {
        dispatch(setScreenSignup({ visible: false }));
      },
    }),
    [dispatch],
  );
}
