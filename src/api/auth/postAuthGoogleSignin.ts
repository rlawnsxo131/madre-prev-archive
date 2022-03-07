import apiClient from '../apiClient';

interface PostAuthGoogleSigninParams {
  accessToken: string;
}

interface PostAuthGoogleSigninResponse {}

export default async function postAuthGoogleSignin({
  accessToken,
}: PostAuthGoogleSigninParams) {
  const { data } = await apiClient.post<PostAuthGoogleSigninResponse | any>(
    `/auth/google/signin`,
    {
      access_token: accessToken,
    },
  );
  return data;
}
