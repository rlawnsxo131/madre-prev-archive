import { createSlice } from '@reduxjs/toolkit';

interface HomeState {
  header: {
    navigation: {
      visible: boolean;
    };
  };
}

const initialState: HomeState = {
  header: {
    navigation: {
      visible: false,
    },
  },
};

const homeSlice = createSlice({
  name: 'home',
  initialState,
  reducers: {
    setHomeMobileNavigation(state) {
      const { visible } = state.header.navigation;
      state.header.navigation.visible = !visible;
    },
  },
});

export const { setHomeMobileNavigation } = homeSlice.actions;

export default homeSlice.reducer;
