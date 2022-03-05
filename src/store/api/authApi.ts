import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import postAuthGoogleSignin from '../../api/auth/postAuthGoogleSignin';
import {
  setLoading,
  setPopupAuth,
  setPopupCommon,
  setScreenSignup,
} from '../core';

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
    postGoogleSignin: build.mutation<
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
        try {
          // check google registered
          dispatch(
            setLoading({
              visible: true,
            }),
          );
          await queryFulfilled;

          // excute signin or sinup action
          const { data } = getCacheEntry();
          if (data?.exist) {
            const {} = await postAuthGoogleSignin({ accessToken });
          }
          if (!data?.exist) {
            dispatch(setScreenSignup({ visible: true }));
          }
          dispatch(setPopupAuth({ visible: false }));
          dispatch(
            setLoading({
              visible: false,
            }),
          );
        } catch (e) {
          dispatch(
            setLoading({
              visible: false,
            }),
          );
          dispatch(setPopupAuth({ visible: false }));
          dispatch(
            setPopupCommon({
              title: '에러 발생',
              message: '잠시후 다시 시도해 주세요.',
              visible: true,
            }),
          );
        }
      },
    }),
  }),
});

export const { usePostGoogleSigninMutation } = authApi;

export default authApi;
