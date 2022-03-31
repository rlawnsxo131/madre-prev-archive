import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import {
  PostAuthGoogleSigninParams,
  PostAuthGoogleSigninResponse,
  PostAuthGoogleSignupParams,
  PostAuthGoogleSignupResponse,
} from '../../@types/api/auth.type';
import postAuthGoogleSignin from '../../api/auth/postAuthGoogleSignin';
import common from '../common';
import popupAuth from '../popupAuth';
import screenSignup from '../screenSignup';

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
    postGoogleCheckWithSignin: build.mutation<
      PostAuthGoogleSigninResponse,
      PostAuthGoogleSigninParams
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
          dispatch(common.actions.showLoading());
          await queryFulfilled;

          // excute signin or sinup action
          const { data } = getCacheEntry();
          if (data?.exist) {
            const data = await postAuthGoogleSignin({ accessToken });
            console.log('signin: ', data);
          } else {
            dispatch(screenSignup.actions.show());
            dispatch(
              screenSignup.actions.setAccessToken({
                accessToken,
              }),
            );
          }
          dispatch(popupAuth.actions.close());
          dispatch(common.actions.closeLoading());
        } catch (e) {
          dispatch(common.actions.closeLoading());
          dispatch(popupAuth.actions.setIsError());
        }
      },
    }),
    postGoogleSignup: build.mutation<
      PostAuthGoogleSignupResponse,
      PostAuthGoogleSignupParams
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
        try {
          dispatch(common.actions.showLoading());
          await queryFulfilled;

          const { data } = getCacheEntry();
          console.log('signup', data);
          dispatch(common.actions.closeLoading());
          dispatch(screenSignup.actions.close());
        } catch (e) {
          dispatch(common.actions.closeLoading());
        }
      },
    }),
  }),
});

export default authApi;
