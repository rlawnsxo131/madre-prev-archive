import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { UserTokenProfile } from '../@types/domain/auth.types';
import { MADRE_USER } from '../constants';
import { Storage } from '../lib/storage';

interface UserState {
  isPending: boolean;
  userTokenProfile: UserTokenProfile | null;
}

const initialState: UserState = {
  isPending: true,
  userTokenProfile: null,
};

const user = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser(state, action: PayloadAction<Pick<UserState, 'userTokenProfile'>>) {
      Storage.setItem(MADRE_USER, action.payload.userTokenProfile);
      state.userTokenProfile = action.payload.userTokenProfile;
    },
    resetUser(state) {
      Storage.removeItem(MADRE_USER);
      state.userTokenProfile = null;
    },
    setIsPending(state, action: PayloadAction<Pick<UserState, 'isPending'>>) {
      state.isPending = action.payload.isPending;
    },
  },
});

export default user;
