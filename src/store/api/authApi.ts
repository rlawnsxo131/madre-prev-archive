import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

/**
 * https://redux-toolkit.js.org/rtk-query/usage-with-typescript
 */
const authApi = createApi({
  reducerPath: 'authApi',
  baseQuery: fetchBaseQuery({
    baseUrl: `${process.env.REACT_APP_API_URI}/api/v1/auth`,
    prepareHeaders(headers) {
      return headers;
    },
    credentials: 'include',
  }),
  endpoints: (build) => ({
    getAuthCheckGoogle: build.query<string, any>({
      query: () => '/google',
      async onQueryStarted(_, { queryFulfilled }) {
        console.log('started');
        await queryFulfilled;
        console.log('end');
      },
    }),
    // currently there is an error in this part
    postAuthCheckGoogle: build.mutation<
      { exist: boolean },
      { accessToken: string }
    >({
      query: ({ ...accessToken }) => ({
        url: '/google/check',
        mehtod: 'POST',
        body: accessToken,
      }),
      async onQueryStarted({ accessToken }, { dispatch, queryFulfilled }) {
        await queryFulfilled;
        console.log(accessToken);
      },
    }),
  }),
});

export const { useGetAuthCheckGoogleQuery, usePostAuthCheckGoogleMutation } =
  authApi;

export default authApi;
