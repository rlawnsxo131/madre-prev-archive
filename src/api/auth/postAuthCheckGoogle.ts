import apiClient from '../apiClient';

interface PostAuthCheckGoogleParams {
  accessToken: string;
}

interface PostAuthCheckGoogleResult {
  exists: boolean;
}

export default async function postAuthCheckGoogle({
  accessToken,
}: PostAuthCheckGoogleParams) {
  const response = await apiClient.post<PostAuthCheckGoogleResult>(
    `/auth/google/check`,
    {
      access_token: accessToken,
    },
  );
  return response.data.exists;
}
