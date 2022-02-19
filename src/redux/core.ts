import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface PopupCommonState {
  visible: boolean;
  title: string;
  message: string;
}

interface PopupLoginState {
  visible: boolean;
}

interface CoreState {
  popupCommon: PopupCommonState;
  popupLogin: PopupLoginState;
}

const initialState: CoreState = {
  popupCommon: {
    visible: false,
    title: '',
    message: '',
  },
  popupLogin: {
    visible: false,
  },
};

const coreSlice = createSlice({
  name: 'core',
  initialState,
  reducers: {
    setPopupCommon(state, action: PayloadAction<PopupCommonState>) {
      const { visible, title, message } = action.payload;
      state.popupCommon.visible = visible;
      state.popupCommon.title = title;
      state.popupCommon.message = message;
    },
    setPopupLogin(state, action: PayloadAction<PopupLoginState>) {
      const { visible } = action.payload;
      state.popupLogin.visible = visible;
    },
  },
});

export const { setPopupCommon, setPopupLogin } = coreSlice.actions;

export default coreSlice.reducer;
