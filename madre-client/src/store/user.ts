import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { UserProfile } from '../@types/domain/account.types';
import { MADRE_USER_PROFILE } from '../constants';
import { Storage } from '../lib/storage';

interface UserLoadStatus {
  isPending: boolean;
  isError: boolean;
}

interface UserState {
  menu: {
    visible: boolean;
  };
  loadUserStatus: UserLoadStatus;
  profile: UserProfile | null;
}

const initialState: UserState = {
  menu: {
    visible: false,
  },
  loadUserStatus: {
    isPending: false,
    isError: false,
  },
  profile: null,
};

const user = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser(state, action: PayloadAction<Pick<UserState, 'profile'>>) {
      Storage.setItem(MADRE_USER_PROFILE, action.payload.profile);
      state.profile = action.payload.profile;
    },
    resetUser(state) {
      Storage.removeItem(MADRE_USER_PROFILE);
      state.profile = null;
    },
    setLoadUserStatusIsPending(
      state,
      action: PayloadAction<Pick<UserLoadStatus, 'isPending'>>,
    ) {
      state.loadUserStatus.isPending = action.payload.isPending;
    },
    setLoadUserStatusIsError(
      state,
      action: PayloadAction<Pick<UserLoadStatus, 'isError'>>,
    ) {
      state.loadUserStatus.isError = action.payload.isError;
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
