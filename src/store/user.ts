import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface UserState {
  token: string;
}

const initialState: UserState = {
  token: '',
};

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser(state, action: PayloadAction<UserState>) {
      const { token } = action.payload;
      state.token = token;
    },
  },
});

export default userSlice.reducer;
