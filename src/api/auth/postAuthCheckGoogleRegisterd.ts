import apiClient from '../apiClient';

interface PostAuthCheckGoogleRegisterdParams {
  accessToken: string;
}

interface PostAuthCheckGoogleRegisterdResult {
  exists: boolean;
}

export default async function postAuthCheckGoogleRegisterd({
  accessToken,
}: PostAuthCheckGoogleRegisterdParams) {
  const response = await apiClient.post<PostAuthCheckGoogleRegisterdResult>(
    `/auth/google/check/registerd`,
    {
      access_token: accessToken,
    },
  );
  return response.data.exists;
}
