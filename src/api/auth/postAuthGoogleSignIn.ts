import {
  PostAuthGoogleSignInParams,
  PostAuthGoogleSigInResponse,
} from '../../@types/api/auth.type';
import apiClient from '../apiClient';

export default async function postAuthGoogleSignIn({
  access_token,
}: PostAuthGoogleSignInParams) {
  const { data } = await apiClient.post<PostAuthGoogleSigInResponse>(
    `/auth/google/sign-in`,
    {
      access_token: access_token,
    },
  );
  return data;
}
