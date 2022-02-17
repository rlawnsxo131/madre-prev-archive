import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface HomeState {
  mobileNavigationState: {
    visible: boolean;
  };
}

const initialState: HomeState = {
  mobileNavigationState: {
    visible: false,
  },
};

const homeSlice = createSlice({
  name: 'home',
  initialState,
  reducers: {
    setMobileNavigation(state, action: PayloadAction) {
      const { visible } = state.mobileNavigationState;
      state.mobileNavigationState.visible = !visible;
    },
  },
});

export const { setMobileNavigation } = homeSlice.actions;

export default homeSlice.reducer;
