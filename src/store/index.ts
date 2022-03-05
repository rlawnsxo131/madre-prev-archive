import { combineReducers, configureStore } from '@reduxjs/toolkit';
import authApi from './api/authApi';
import common from './common';
import home from './home';
import popupAuth from './popupAuth';
import screenSignup from './screenSignup';
import theme from './theme';
import user from './user';

const rootReducer = combineReducers({
  common,
  home,
  theme,
  user,
  // in core component
  popupAuth,
  screenSignup,
  // rtk queries
  [authApi.reducerPath]: authApi.reducer,
});

export const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(authApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
