import { useEffect, useRef } from 'react';
import authApi from '../../../../store/api/authApi';

/**
 * https://developers.google.com/identity/protocols/oauth2/scopes
 * ex)
 * scope:
 * 'https://www.googleapis.com/auth/calendar.readonly \
 *  https://www.googleapis.com/auth/contacts.readonly'
 */
export default function useButtonGoogleSignIn() {
  const buttonRef = useRef<HTMLButtonElement>(null);
  const [googleCheckWithSignIn] =
    authApi.usePostGoogleCheckWithSignInMutation();

  useEffect(() => {
    if (!buttonRef.current) return;
    const client = window.google.accounts.oauth2.initTokenClient({
      client_id:
        '939741412461-vcst6mvh6s3mv9qcgrnp91p48bf62gdi.apps.googleusercontent.com',
      scope: 'https://www.googleapis.com/auth/userinfo.profile',
      callback: async (tokenResponse: any) => {
        const access_token = tokenResponse.access_token;
        await googleCheckWithSignIn({ access_token });
      },
    });
    buttonRef.current.onclick = () => client.requestAccessToken();
  }, [googleCheckWithSignIn]);

  return {
    buttonRef,
  };
}
