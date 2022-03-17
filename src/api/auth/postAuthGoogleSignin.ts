import apiClient from '../apiClient';

export interface PostAuthGoogleSigninParams {
  accessToken: string;
}

export interface PostAuthGoogleSigninResponse {
  exist: boolean;
}

export default async function postAuthGoogleSignin({
  accessToken,
}: PostAuthGoogleSigninParams) {
  const { data } = await apiClient.post<PostAuthGoogleSigninResponse>(
    `/auth/google/signin`,
    {
      access_token: accessToken,
    },
  );
  return data;
}
