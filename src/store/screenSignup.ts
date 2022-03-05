import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface ScreenSignupState {
  visible: boolean;
}

const initialState: ScreenSignupState = {
  visible: false,
};

const screenSignupSlice = createSlice({
  name: 'screenSignup',
  initialState,
  reducers: {
    setScreenSignup(state, action: PayloadAction<ScreenSignupState>) {
      const { visible } = action.payload;
      state.visible = visible;
    },
  },
});

export const { setScreenSignup } = screenSignupSlice.actions;

export default screenSignupSlice.reducer;
