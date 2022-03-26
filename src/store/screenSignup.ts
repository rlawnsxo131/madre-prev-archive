import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface ScreenSignupState {
  visible: boolean;
  accessToken: string;
  username: string;
}

const initialState: ScreenSignupState = {
  visible: false,
  accessToken: '',
  username: '',
};

const screenSignup = createSlice({
  name: 'screenSignup',
  initialState,
  reducers: {
    show(state) {
      state.visible = true;
    },
    close(state) {
      state.visible = false;
      state.accessToken = '';
      state.username = '';
    },
    setAccessToken(
      state,
      action: PayloadAction<Pick<ScreenSignupState, 'accessToken'>>,
    ) {
      state.accessToken = action.payload.accessToken;
    },
    setUsername(
      state,
      action: PayloadAction<Pick<ScreenSignupState, 'username'>>,
    ) {
      state.username = action.payload.username;
    },
  },
});

export default screenSignup;
