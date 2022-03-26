import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface PopupCommonState {
  visible: boolean;
  title: string;
  message: string;
}

interface LoadingState {
  visible: boolean;
}

interface CommonState {
  popupCommon: PopupCommonState;
  loading: LoadingState;
}

const initialState: CommonState = {
  popupCommon: {
    visible: false,
    title: '',
    message: '',
  },
  loading: {
    visible: false,
  },
};

const common = createSlice({
  name: 'common',
  initialState,
  reducers: {
    setPopupCommon(state, action: PayloadAction<PopupCommonState>) {
      const { visible, title, message } = action.payload;
      state.popupCommon.visible = visible;
      state.popupCommon.title = title;
      state.popupCommon.message = message;
    },
    setLoading(state, action: PayloadAction<LoadingState>) {
      state.loading.visible = action.payload.visible;
    },
  },
});

export default common;
