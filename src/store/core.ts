import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface PopupCommonState {
  visible: boolean;
  title: string;
  message: string;
}

interface PopupAuthState {
  visible: boolean;
}

interface ScreenSignupState {
  visible: boolean;
}

interface LoadingState {
  visible: boolean;
}

interface CoreState {
  popupCommon: PopupCommonState;
  popupAuth: PopupAuthState;
  screenSignup: ScreenSignupState;
  loading: LoadingState;
}

const initialState: CoreState = {
  popupCommon: {
    visible: false,
    title: '',
    message: '',
  },
  popupAuth: {
    visible: false,
  },
  screenSignup: {
    visible: false,
  },
  loading: {
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
    setPopupAuth(state, action: PayloadAction<PopupAuthState>) {
      const { visible } = action.payload;
      state.popupAuth.visible = visible;
    },
    setScreenSignup(state, action: PayloadAction<ScreenSignupState>) {
      const { visible } = action.payload;
      state.screenSignup.visible = visible;
    },
    setLoading(state, action: PayloadAction<LoadingState>) {
      const { visible } = action.payload;
      state.loading.visible = visible;
    },
  },
});

export const { setPopupCommon, setPopupAuth, setScreenSignup, setLoading } =
  coreSlice.actions;

export default coreSlice.reducer;
