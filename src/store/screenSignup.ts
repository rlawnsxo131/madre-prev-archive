import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface ScreenSignupState {
  visible: boolean;
}

const initialState: ScreenSignupState = {
  visible: false,
};

const screenSignup = createSlice({
  name: 'screenSignup',
  initialState,
  reducers: {
    setVisible(state, action: PayloadAction<ScreenSignupState>) {
      const { visible } = action.payload;
      state.visible = visible;
    },
  },
});

export default screenSignup;
