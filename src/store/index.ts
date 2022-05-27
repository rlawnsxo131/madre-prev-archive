import { combineReducers, configureStore } from '@reduxjs/toolkit';
import authApi from './api/authApi';
import common from './common';
import popupAuth from './popupAuth';
import screenSignUp from './screenSignUp';
import theme from './theme';
import user from './user';

const rootReducer = combineReducers({
  common: common.reducer,
  theme: theme.reducer,
  user: user.reducer,
  popupAuth: popupAuth.reducer,
  screenSignUp: screenSignUp.reducer,
  [authApi.reducerPath]: authApi.reducer,
});

export const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(authApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
