import apiClient from '../apiClient';

interface PostAuthGoogleCheckParams {
  accessToken: string;
}

interface PostAuthGoogleCheckResponse {
  exist: boolean;
}

export default async function postAuthGoogleCheck({
  accessToken,
}: PostAuthGoogleCheckParams) {
  const { data } = await apiClient.post<PostAuthGoogleCheckResponse>(
    `/auth/google/check`,
    {
      access_token: accessToken,
    },
  );
  return data;
}
