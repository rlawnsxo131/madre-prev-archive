import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { UserTokenProfile } from '../@types/domain/auth.types';
import { MADRE_USER } from '../constants';
import { Storage } from '../lib/storage';

interface UserState {
  isPending: boolean;
  userProfile: UserTokenProfile | null;
}

const initialState: UserState = {
  isPending: true,
  userProfile: null,
};

const user = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser(state, action: PayloadAction<Pick<UserState, 'userProfile'>>) {
      Storage.setItem(MADRE_USER, action.payload.userProfile);
      state.userProfile = action.payload.userProfile;
    },
    resetUser(state) {
      Storage.removeItem(MADRE_USER);
      state.userProfile = null;
    },
    setIsPending(state, action: PayloadAction<Pick<UserState, 'isPending'>>) {
      state.isPending = action.payload.isPending;
    },
  },
});

export default user;
