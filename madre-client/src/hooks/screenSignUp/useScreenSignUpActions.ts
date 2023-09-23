import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import screenSignUp from '../../store/screenSignUp';

export default function useScreenSignUpActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      close() {
        dispatch(screenSignUp.actions.close());
      },
      setIsValidateError() {
        dispatch(screenSignUp.actions.setIsValidateError());
      },
      resetIsValidateError() {
        dispatch(screenSignUp.actions.resetIsValidateError());
      },
      resetIsConflictError() {
        dispatch(screenSignUp.actions.resetIsConflictError());
      },
    }),
    [dispatch],
  );
}
