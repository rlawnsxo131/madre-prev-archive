import apiClient from '../apiClient';

interface PostAuthGoogleCheckParams {
  accessToken: string;
}

interface PostAuthGoogleCheckResult {
  exist: boolean;
}

export default async function postAuthGoogleCheck({
  accessToken,
}: PostAuthGoogleCheckParams) {
  const { data } = await apiClient.post<PostAuthGoogleCheckResult>(
    `/auth/google/check`,
    {
      access_token: accessToken,
    },
  );
  return data;
}
