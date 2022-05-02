import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { UserTokenProfile } from '../@types/domain/auth.types';
import { MADRE_USER_TOKEN_PROFILE } from '../constants';
import { Storage } from '../lib/storage';

interface UserState {
  isPending: boolean;
  menu: {
    navigationVisible: boolean;
  };
  userTokenProfile: UserTokenProfile | null;
}

const initialState: UserState = {
  isPending: true,
  menu: {
    navigationVisible: false,
  },
  userTokenProfile: null,
};

const user = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser(state, action: PayloadAction<Pick<UserState, 'userTokenProfile'>>) {
      Storage.setItem(
        MADRE_USER_TOKEN_PROFILE,
        action.payload.userTokenProfile,
      );
      state.userTokenProfile = action.payload.userTokenProfile;
    },
    resetUser(state) {
      Storage.removeItem(MADRE_USER_TOKEN_PROFILE);
      state.userTokenProfile = null;
    },
    setIsPending(state, action: PayloadAction<Pick<UserState, 'isPending'>>) {
      state.isPending = action.payload.isPending;
    },
    handlePersonalMenuNavigation(state) {
      const visible = !state.menu.navigationVisible;
      state.menu.navigationVisible = visible;
    },
  },
});

export default user;
