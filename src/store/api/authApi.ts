import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import postAuthGoogleSignin from '../../api/auth/postAuthGoogleSignin';
import { setLoading } from '../common';
import { setPopupAuthIsError, setPopupAuthVisible } from '../popupAuth';
import { setScreenSignup } from '../screenSignup';

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
            const data = await postAuthGoogleSignin({ accessToken });
            console.log(data);
          }
          if (!data?.exist) {
            dispatch(setScreenSignup({ visible: true }));
          }
          dispatch(setPopupAuthVisible({ visible: false }));
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
          dispatch(setPopupAuthIsError({ isError: true }));
        }
      },
    }),
    postGoogleSignup: build.mutation<
      { access_token: string; username: string },
      any
    >({
      query: ({ accessToken, username }) => ({
        url: '/google/signup',
        method: 'POST',
        body: {
          access_token: accessToken,
          username,
        },
      }),
      async onQueryStarted(_, { dispatch, queryFulfilled, getCacheEntry }) {
        dispatch(setLoading({ visible: true }));
        await queryFulfilled;

        const { data } = getCacheEntry();
        console.log(data);
      },
    }),
  }),
});

export const { usePostGoogleSigninMutation, usePostGoogleSignupMutation } =
  authApi;

export default authApi;
