import apiClient from '../apiClient';

interface PostAuthGoogleSignupParams {
  accessToken: string;
}

interface PostAuthGoogleSignupResponse {}

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
