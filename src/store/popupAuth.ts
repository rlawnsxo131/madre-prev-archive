import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface PopupAuthState {
  visible: boolean;
  isError: boolean;
}

const initialState: PopupAuthState = {
  visible: false,
  isError: false,
};

const popupAuthSlice = createSlice({
  name: 'popupAuth',
  initialState,
  reducers: {
    setPopupAuthVisible(state, action: PayloadAction<{ visible: boolean }>) {
      const { visible } = action.payload;
      state.visible = visible;
    },
    setPopupAuthIsError(state, action: PayloadAction<{ isError: boolean }>) {
      const { isError } = action.payload;
      state.isError = isError;
    },
  },
});

export const { setPopupAuthVisible, setPopupAuthIsError } =
  popupAuthSlice.actions;

export default popupAuthSlice.reducer;
