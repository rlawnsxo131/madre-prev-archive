import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface PopupCommonState {
  visible: boolean;
  title: string;
  message: string;
}

interface CoreState {
  popupCommon: PopupCommonState;
}

const initialState: CoreState = {
  popupCommon: {
    visible: false,
    title: '',
    message: '',
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
  },
});

export const { setPopupCommon } = coreSlice.actions;

export default coreSlice.reducer;
