import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface UserState {
  access_token: string;
  display_name: string;
}

const initialState: UserState = {
  access_token: '',
  display_name: ''
};

const user = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser(state, action: PayloadAction<UserState>) {
      const {access_token, display_name} = action.payload;
      state.access_token = access_token;
      state.display_name = display_name;
    },
  },
});

export default user;
