import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface PopupAuthState {
  visible: boolean;
  isError: boolean;
}

const initialState: PopupAuthState = {
  visible: false,
  isError: false,
};

const popupAuth = createSlice({
  name: 'popupAuth',
  initialState,
  reducers: {
    setVisible(state, action: PayloadAction<{ visible: boolean }>) {
      const { visible } = action.payload;
      state.visible = visible;
    },
    setIsError(state, action: PayloadAction<{ isError: boolean }>) {
      const { isError } = action.payload;
      state.isError = isError;
    },
  },
});

export default popupAuth;
