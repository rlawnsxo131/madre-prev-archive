import { createSlice } from '@reduxjs/toolkit';

interface UserState {
  token: string;
}

const initialState: UserState = {
  token: '',
};

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {},
});

export default userSlice.reducer;
