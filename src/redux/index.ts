import { combineReducers, configureStore } from '@reduxjs/toolkit';
import authApi from './api/authApi';
import core from './core';
import home from './home';
import theme from './theme';

const rootReducer = combineReducers({
  home,
  theme,
  core,
  [authApi.reducerPath]: authApi.reducer,
});

export const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(authApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
