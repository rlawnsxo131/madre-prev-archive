import {
  PostAuthGoogleSigninParams,
  PostAuthGoogleSigninResponse,
} from '../../@types/api/auth.type';
import apiClient from '../apiClient';

export default async function postAuthGoogleSignin({
  accessToken,
}: PostAuthGoogleSigninParams) {
  const { data } = await apiClient.post<PostAuthGoogleSigninResponse>(
    `/auth/google/signin`,
    {
      accessToken,
    },
  );
  return data;
}
