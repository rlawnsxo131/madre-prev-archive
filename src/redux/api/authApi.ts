import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

/**
 * https://redux-toolkit.js.org/rtk-query/usage-with-typescript
 */
const authApi = createApi({
  reducerPath: 'authApi',
  baseQuery: fetchBaseQuery({
    baseUrl: `${process.env.REACT_APP_API_URI}/api/v1/auth/`,
    prepareHeaders(headers) {
      return headers;
    },
    credentials: 'include',
  }),
  endpoints: (build) => ({
    getAuthCheckGoogle: build.query<string, any>({
      query: () => 'google',
    }),
  }),
});

export const { useGetAuthCheckGoogleQuery } = authApi;

export default authApi;
