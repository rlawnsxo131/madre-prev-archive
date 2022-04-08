import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import {
  PostAuthGoogleCheckParams,
  PostAuthGoogleCheckResponse,
  PostAuthGoogleSignUpParams,
  PostAuthGoogleSignUpResponse,
} from '../../@types/api/auth.type';
import postAuthGoogleSignIn from '../../api/auth/postAuthGoogleSignIn';
import common from '../common';
import popupAuth from '../popupAuth';
import screenSignUp from '../screenSignUp';
import user from '../user';

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
    postGoogleCheckWithSignIn: build.mutation<
      PostAuthGoogleCheckResponse,
      PostAuthGoogleCheckParams
    >({
      query: ({ access_token }) => ({
        url: '/google/check',
        method: 'POST',
        body: {
          access_token,
        },
      }),
      async onQueryStarted(
        { access_token },
        { dispatch, queryFulfilled, getCacheEntry },
      ) {
        try {
          // check google registered
          dispatch(common.actions.showLoading());
          await queryFulfilled;

          // excute sign-in or sign-up action
          const { data } = getCacheEntry();
          if (data?.exist) {
            const data = await postAuthGoogleSignIn({ access_token });
            console.log('signin: ', data);
          } else {
            dispatch(screenSignUp.actions.show());
            dispatch(
              screenSignUp.actions.setAccessToken({
                access_token,
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
    postGoogleSignUp: build.mutation<
      PostAuthGoogleSignUpResponse,
      PostAuthGoogleSignUpParams
    >({
      query: ({ access_token, display_name }) => ({
        url: '/google/sign-up',
        method: 'POST',
        body: {
          access_token,
          display_name,
        },
      }),
      async onQueryStarted(_, { dispatch, queryFulfilled, getCacheEntry }) {
        try {
          dispatch(common.actions.showLoading());
          await queryFulfilled;

          const { data } = getCacheEntry();
          if (!data) {
            dispatch(screenSignUp.actions.setIsError());
          } else {
            dispatch(
              user.actions.setUser({
                access_token: data.access_token,
                display_name: data.display_name,
              }),
            );
          }
          dispatch(common.actions.closeLoading());
          dispatch(screenSignUp.actions.close());
        } catch (e) {
          dispatch(common.actions.closeLoading());
          dispatch(screenSignUp.actions.setIsError());
        }
      },
    }),
  }),
});

export default authApi;
