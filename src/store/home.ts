import { createSlice } from '@reduxjs/toolkit';

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
    setHomeMobileNavigation(state) {
      const { visible } = state.mobileNavigationState;
      state.mobileNavigationState.visible = !visible;
    },
  },
});

export const { setHomeMobileNavigation } = homeSlice.actions;

export default homeSlice.reducer;
