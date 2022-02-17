import { combineReducers, configureStore } from '@reduxjs/toolkit';
import core from './core';
import home from './home';
import theme from './theme';

const rootReducer = combineReducers({
  home,
  theme,
  core,
});

export const store = configureStore({
  reducer: rootReducer,
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
