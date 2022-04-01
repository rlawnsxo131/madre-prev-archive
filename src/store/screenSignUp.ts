import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface ScreenSignUpState {
  visible: boolean;
  accessToken: string;
  username: string;
}

const initialState: ScreenSignUpState = {
  visible: false,
  accessToken: '',
  username: '',
};

const screenSignUp = createSlice({
  name: 'screenSignUp',
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
      action: PayloadAction<Pick<ScreenSignUpState, 'accessToken'>>,
    ) {
      state.accessToken = action.payload.accessToken;
    },
    setUsername(
      state,
      action: PayloadAction<Pick<ScreenSignUpState, 'username'>>,
    ) {
      state.username = action.payload.username;
    },
  },
});

export default screenSignUp;
