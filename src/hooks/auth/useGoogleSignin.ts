import { useCallback } from 'react';
import { useDispatch } from 'react-redux';
import postAuthGoogleCheck from '../../api/auth/postAuthGoogleCheck';
import postAuthGoogleSignin from '../../api/auth/postAuthGoogleSignin';
import { AppDispatch } from '../../store';
import useLoadingActions from '../core/useLoadingActions';
import usePopupAuthActions from '../core/usePopupAuthActions';
import useScreenSignupActions from '../core/useScreenSignupActions';

export default function useGoogleSignin() {
  const dispatch = useDispatch<AppDispatch>();
  const { close: closePopupAuth } = usePopupAuthActions();
  const { show: showSignupScreen } = useScreenSignupActions();
  const { show: showLoading, close: closeLoading } = useLoadingActions();

  const googleSignin = useCallback(
    async (accessToken: string) => {
      try {
        showLoading();
        const { exist } = await postAuthGoogleCheck({ accessToken });
        closePopupAuth();
        if (exist) {
          const {} = await postAuthGoogleSignin({ accessToken });
        } else {
          showSignupScreen();
        }
      } catch (e) {
        console.log(e);
      } finally {
        closeLoading();
      }
    },
    [dispatch],
  );

  return {
    googleSignin,
  };
}
