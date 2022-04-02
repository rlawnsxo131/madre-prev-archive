import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import screenSignUp from '../../store/screenSignUp';

export default function useScreenSignUpActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      show() {
        dispatch(screenSignUp.actions.show());
      },
      close() {
        dispatch(screenSignUp.actions.close());
      },
      setError() {
        dispatch(screenSignUp.actions.setIsError());
      },
      resetError() {
        dispatch(screenSignUp.actions.resetIsError());
      },
    }),
    [dispatch],
  );
}
