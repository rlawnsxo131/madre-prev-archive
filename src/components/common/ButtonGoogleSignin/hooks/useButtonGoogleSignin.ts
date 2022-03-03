import { useEffect, useRef } from 'react';
import postAuthGoogleCheck from '../../../../api/auth/postAuthGoogleCheck';
import useLoadingActions from '../../../../hooks/core/useLoadingActions';
import usePopupAuthActions from '../../../../hooks/core/usePopupAuthActions';
import useScreenSignupActions from '../../../../hooks/core/useScreenSignupActions';
import { usePostAuthCheckGoogleMutation } from '../../../../store/api/authApi';

export default function useButtonGoogleSignin() {
  const buttonRef = useRef<HTMLButtonElement>(null);
  const { onClose: onClosePopupAuth } = usePopupAuthActions();
  const { onShow: onShowSignupScreen } = useScreenSignupActions();
  const { onShow: onShowLoading, onClose: onCloseLoading } =
    useLoadingActions();
  // const [postAuthCheckGoogle, exist] = usePostAuthCheckGoogleMutation();

  useEffect(() => {
    if (!buttonRef.current) return;
    if (!window.gapi) return;
    window.gapi.load('auth2', function () {
      // Retrieve the singleton for the GoogleAuth library and set up the client.
      const auth2 = window.gapi.auth2.init({
        client_id:
          '939741412461-vcst6mvh6s3mv9qcgrnp91p48bf62gdi.apps.googleusercontent.com',
        cookiepolicy: 'single_host_origin',
      });

      auth2.attachClickHandler(
        buttonRef.current,
        {},
        async (googleUser: any) => {
          const accessToken = googleUser?.getAuthResponse(true).access_token;
          // postAuthCheckGoogle({ accessToken });

          onShowLoading();
          const { exist } = await postAuthGoogleCheck({
            accessToken,
          });
          onCloseLoading();
          if (exist) {
            // signin
            return;
          }
          onClosePopupAuth();
          onShowSignupScreen();
        },
      );
    });
  }, [window.gapi, buttonRef.current]);

  return {
    buttonRef,
  };
}
