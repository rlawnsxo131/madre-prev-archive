import apiClient from '../apiClient';

export interface PostAuthGoogleSignupParams {
  accessToken: string;
  username: string;
}

export interface PostAuthGoogleSignupResponse {}

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
