import {
  PostAuthGoogleSignupParams,
  PostAuthGoogleSignupResponse,
} from '../../@types/api/auth.type';
import apiClient from '../apiClient';

export default async function postAuthGoogleSignup({
  accessToken,
}: PostAuthGoogleSignupParams) {
  const { data } = await apiClient.post<PostAuthGoogleSignupResponse>(
    `/auth/google/signup`,
    {
      access_token: accessToken,
    },
  );
  return data;
}
