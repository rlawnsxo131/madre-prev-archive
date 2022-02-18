import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

const authApi = createApi({
  reducerPath: 'authApi',
  baseQuery: fetchBaseQuery({
    baseUrl: `${process.env.REACT_APP_API_URI}/api/v1/auth/`,
  }),
  endpoints: (builder) => ({
    getAuthCheckGoogle: builder.query<string, any>({
      query: () => 'google',
    }),
  }),
});

export const { useGetAuthCheckGoogleQuery } = authApi;

export default authApi;
