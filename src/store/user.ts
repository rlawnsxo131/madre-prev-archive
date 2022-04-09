import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { MADRE_USER } from '../constants';
import { Storage } from '../lib/storage';

interface UserState {
  isPending: boolean;
  access_token: string;
  display_name: string;
}

const initialState: UserState = {
  isPending: true,
  access_token: '',
  display_name: '',
};

const user = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser(
      state,
      action: PayloadAction<Pick<UserState, 'access_token' | 'display_name'>>,
    ) {
      const { access_token, display_name } = action.payload;
      Storage.setItem(MADRE_USER, {
        access_token,
        display_name,
      });
      state.access_token = access_token;
      state.display_name = display_name;
    },
    setIsPending(state, action: PayloadAction<Pick<UserState, 'isPending'>>) {
      state.isPending = action.payload.isPending;
    },
  },
});

export default user;
