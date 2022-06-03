import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { ResponseError } from '../../@types/api/api.types';
import {
  GetAuthResponse,
  PostAuthGoogleCheckParams,
  PostAuthGoogleCheckResponse,
  PostAuthGoogleSignUpParams,
  PostAuthGoogleSignUpResponse,
} from '../../@types/api/auth.types';
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
    referrerPolicy: 'strict-origin-when-cross-origin',
  }),
  // refetchOnFocus: true,
  tagTypes: ['Auth'],
  endpoints: (build) => ({
    get: build.query<GetAuthResponse, undefined>({
      query() {
        return {
          url: '',
        };
      },
      async onQueryStarted(_, { dispatch, queryFulfilled, getCacheEntry }) {
        try {
          await queryFulfilled;

          const { data } = getCacheEntry();
          dispatch(
            user.actions.setLoadUserStatusIsError({
              isError: false,
            }),
          );
          if (data) {
            dispatch(
              user.actions.setUser({
                profile: data,
              }),
            );
          } else {
            dispatch(user.actions.resetUser());
          }
        } catch (e) {
          dispatch(
            user.actions.setLoadUserStatusIsError({
              isError: true,
            }),
          );
          console.log('error: ', e);
        }
      },
      providesTags: ['Auth'],
    }),
    delete: build.mutation<undefined, undefined>({
      query() {
        return {
          url: '',
          method: 'DELETE',
        };
      },
      async onQueryStarted(_, { dispatch, queryFulfilled }) {
        try {
          dispatch(common.actions.showLoading());
          await queryFulfilled;

          dispatch(user.actions.resetUser());
          dispatch(user.actions.closeNavigation());
          dispatch(common.actions.closeLoading());
        } catch (e) {
          dispatch(common.actions.closeLoading());
          dispatch(user.actions.closeNavigation());

          // TODO: Let's think about error formalization
          const { error } = e as ResponseError;
          if (error.data.status === 401) {
            dispatch(
              common.actions.showPopupCommon({
                title: '',
                message: '잘못된 요청입니다.',
              }),
            );
          } else {
            dispatch(
              common.actions.showPopupCommon({
                title: '',
                message: '로그아웃 에러가 발생했습니다',
              }),
            );
          }
          console.log(e);
        }
      },
      invalidatesTags: ['Auth'],
    }),
    postGoogleCheckWithSignIn: build.mutation<
      PostAuthGoogleCheckResponse,
      PostAuthGoogleCheckParams
    >({
      query({ access_token }) {
        return {
          url: '/google/check',
          method: 'POST',
          body: {
            access_token,
          },
        };
      },
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
            const data = await postAuthGoogleSignIn({
              access_token,
            });
            dispatch(
              user.actions.setUser({
                profile: data,
              }),
            );
          } else {
            dispatch(
              screenSignUp.actions.show({
                access_token,
              }),
            );
          }
          dispatch(popupAuth.actions.close());
          dispatch(common.actions.closeLoading());
        } catch (e) {
          dispatch(common.actions.closeLoading());
          dispatch(popupAuth.actions.setIsError());
          console.log(e);
        }
      },
    }),
    postGoogleSignUp: build.mutation<
      PostAuthGoogleSignUpResponse,
      PostAuthGoogleSignUpParams
    >({
      query({ access_token, username }) {
        return {
          url: '/google/sign-up',
          method: 'POST',
          body: {
            access_token,
            username,
          },
        };
      },
      async onQueryStarted(_, { dispatch, queryFulfilled, getCacheEntry }) {
        try {
          dispatch(common.actions.showLoading());
          dispatch(screenSignUp.actions.resetIsError());
          dispatch(screenSignUp.actions.resetIsConflictError());
          await queryFulfilled;

          const { data } = getCacheEntry();
          if (!data) {
            dispatch(screenSignUp.actions.setIsError());
          } else {
            dispatch(
              user.actions.setUser({
                profile: data,
              }),
            );
          }
          dispatch(common.actions.closeLoading());
          dispatch(screenSignUp.actions.close());
        } catch (e) {
          dispatch(common.actions.closeLoading());

          const { error } = e as ResponseError;
          if (error.data.status === 400) {
            dispatch(screenSignUp.actions.setIsValidateError());
          } else if (error.data.status === 409) {
            dispatch(screenSignUp.actions.setIsConflictError());
          } else {
            dispatch(screenSignUp.actions.setIsError());
          }
          console.log(e);
        }
      },
    }),
  }),
});

export default authApi;
