import {
  PostAuthGoogleSignUpParams,
  PostAuthGoogleSignUpResponse,
} from '../../@types/api/auth.type';
import apiClient from '../apiClient';

export default async function postAuthGoogleSignUp({
  access_token,
}: PostAuthGoogleSignUpParams) {
  const { data } = await apiClient.post<PostAuthGoogleSignUpResponse>(
    `/auth/google/sign-up`,
    {
      access_token,
    },
  );
  return data;
}
