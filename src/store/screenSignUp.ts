import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface ScreenSignUpState {
  visible: boolean;
  isError: boolean;
  access_token: string;
}

const initialState: ScreenSignUpState = {
  visible: false,
  isError: false,
  access_token: '',
};

const screenSignUp = createSlice({
  name: 'screenSignUp',
  initialState,
  reducers: {
    show(
      state,
      action: PayloadAction<Pick<ScreenSignUpState, 'access_token'>>,
    ) {
      state.visible = true;
      state.access_token = action.payload.access_token;
    },
    close(state) {
      state.visible = false;
      state.isError = false;
      state.access_token = '';
    },
    setIsError(state) {
      state.isError = true;
    },
    resetIsError(state) {
      state.isError = false;
    },
  },
});

export default screenSignUp;
