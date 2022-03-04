import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import postAuthGoogleSignin from '../../api/auth/postAuthGoogleSignin';
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
  tagTypes: ['Auth'],
  endpoints: (build) => ({
    // getAuthCheckGoogle: build.query<string, any>({
    //   query: () => '/google',
    //   async onQueryStarted(_, { queryFulfilled }) {
    //     console.log('started');
    //     await queryFulfilled;
    //     console.log('end');
    //   },
    // }),
    postGoogleSinin: build.mutation<
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
      async onQueryStarted(
        { accessToken },
        { dispatch, queryFulfilled, getCacheEntry },
      ) {
        // check google registered
        dispatch(
          setLoading({
            visible: true,
          }),
        );
        await queryFulfilled;
        dispatch(setPopupAuth({ visible: false }));
        dispatch(
          setLoading({
            visible: false,
          }),
        );

        // excute signin or sinup action
        const { data } = getCacheEntry();
        if (data?.exist) {
          const {} = await postAuthGoogleSignin({ accessToken });
        }
        if (!data?.exist) {
          dispatch(setScreenSignup({ visible: true }));
        }
      },
    }),
  }),
});

export const { usePostGoogleSininMutation } = authApi;

export default authApi;
