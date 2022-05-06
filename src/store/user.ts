import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { UserTokenProfile } from '../@types/domain/auth.types';
import { MADRE_USER_TOKEN_PROFILE } from '../constants';
import { Storage } from '../lib/storage';

interface UserState {
  isPending: boolean;
  menu: {
    visible: boolean;
  };
  profile: UserTokenProfile | null;
}

const initialState: UserState = {
  isPending: true,
  menu: {
    visible: false,
  },
  profile: null,
};

const user = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser(state, action: PayloadAction<Pick<UserState, 'profile'>>) {
      Storage.setItem(MADRE_USER_TOKEN_PROFILE, action.payload.profile);
      state.profile = action.payload.profile;
    },
    resetUser(state) {
      Storage.removeItem(MADRE_USER_TOKEN_PROFILE);
      state.profile = null;
    },
    setIsPending(state, action: PayloadAction<Pick<UserState, 'isPending'>>) {
      state.isPending = action.payload.isPending;
    },
    handleNavigation(state) {
      const visible = !state.menu.visible;
      state.menu.visible = visible;
    },
    closeNavigation(state) {
      state.menu.visible = false;
    },
  },
});

export default user;
