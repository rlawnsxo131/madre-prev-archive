import {
  PostAuthGoogleSignUpParams,
  PostAuthGoogleSignUpResponse,
} from '../../@types/api/auth.type';
import apiClient from '../apiClient';

export default async function postAuthGoogleSignUp({
  accessToken,
}: PostAuthGoogleSignUpParams) {
  const { data } = await apiClient.post<PostAuthGoogleSignUpResponse>(
    `/auth/google/sign-up`,
    {
      access_token: accessToken,
    },
  );
  return data;
}
