import { combineReducers, configureStore } from '@reduxjs/toolkit';
import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux';
import authApi from './api/authApi';
import common from './common';
import popupAuth from './popupAuth';
import screenSignUp from './screenSignUp';
import theme from './theme';
import user from './user';

export const rootReducer = combineReducers({
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
export type AppStore = typeof store;
export type RootReducer = ReturnType<typeof rootReducer>;
export type AppDispatch = AppStore['dispatch'];

export const useAppDispatch: () => AppDispatch = useDispatch;
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
