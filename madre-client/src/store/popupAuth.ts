import { createSlice } from '@reduxjs/toolkit';

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
    show(state) {
      state.visible = true;
    },
    close(state) {
      state.visible = false;
    },
    setIsError(state) {
      state.isError = true;
    },
    resetIsError(state) {
      state.isError = false;
    },
  },
});

export default popupAuth;
