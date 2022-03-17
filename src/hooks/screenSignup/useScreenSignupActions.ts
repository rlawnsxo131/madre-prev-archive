import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import screenSignup from '../../store/screenSignup';

export default function useScreenSignupActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      show() {
        dispatch(screenSignup.actions.setVisible({ visible: true }));
      },
      close() {
        dispatch(screenSignup.actions.setVisible({ visible: false }));
      },
    }),
    [dispatch],
  );
}
