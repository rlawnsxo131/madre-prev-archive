import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { setLoading, setPopupAuth, setScreenSignup } from '../core';

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
  tagTypes: ['User'],
  endpoints: (build) => ({
    // getAuthCheckGoogle: build.query<string, any>({
    //   query: () => '/google',
    //   async onQueryStarted(_, { queryFulfilled }) {
    //     console.log('started');
    //     await queryFulfilled;
    //     console.log('end');
    //   },
    // }),
    postAuthCheckGoogle: build.mutation<
      { exist: boolean },
      { accessToken: string }
    >({
      query: ({ accessToken }) => ({
        url: '/google/check',
        method: 'POST',
        body: {
          access_token: accessToken,
        },
      }),
      async onQueryStarted(_, { dispatch, queryFulfilled, getCacheEntry }) {
        dispatch(
          setLoading({
            visible: true,
          }),
        );
        await queryFulfilled;
        dispatch(setPopupAuth({ visible: false }));
        const { data } = getCacheEntry();
        if (data?.exist) {
          console.log('exist');
        }
        if (!data?.exist) {
          dispatch(setScreenSignup({ visible: true }));
        }
        dispatch(
          setLoading({
            visible: false,
          }),
        );
      },
    }),
  }),
});

export const { usePostAuthCheckGoogleMutation } = authApi;

export default authApi;
