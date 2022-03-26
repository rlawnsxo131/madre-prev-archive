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
    setVisible(state, action: PayloadAction<Pick<PopupAuthState, 'visible'>>) {
      state.visible = action.payload.visible;
    },
    setIsError(state, action: PayloadAction<Pick<PopupAuthState, 'isError'>>) {
      state.isError = action.payload.isError;
    },
  },
});

export default popupAuth;
