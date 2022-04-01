import { useEffect, useRef } from 'react';
import authApi from '../../../../store/api/authApi';

export default function useButtonGoogleSignIn() {
  const buttonRef = useRef<HTMLButtonElement>(null);
  const [googleCheckWithSignIn] =
    authApi.usePostGoogleCheckWithSignInMutation();

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
          await googleCheckWithSignIn({ accessToken });
        },
      );
    });
  }, [window.gapi, buttonRef.current]);

  return {
    buttonRef,
  };
}
