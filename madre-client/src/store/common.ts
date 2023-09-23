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
    showPopupCommon(
      state,
      action: PayloadAction<Pick<PopupCommonState, 'title' | 'message'>>,
    ) {
      const { title, message } = action.payload;
      state.popupCommon.visible = true;
      state.popupCommon.title = title;
      state.popupCommon.message = message;
    },
    closePopupCommon(state) {
      state.popupCommon.visible = false;
      state.popupCommon.title = '';
      state.popupCommon.message = '';
    },
    showLoading(state) {
      state.loading.visible = true;
    },
    closeLoading(state) {
      state.loading.visible = false;
    },
  },
});

export default common;
