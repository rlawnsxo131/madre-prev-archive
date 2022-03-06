import { createSlice, PayloadAction } from '@reduxjs/toolkit';

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
    handleHomeMobileNavigation(state) {
      const { visible } = state.header.navigation;
      state.header.navigation.visible = !visible;
    },
    setHomeMobileNavigation(
      state,
      action: PayloadAction<{ visible: boolean }>,
    ) {
      const { visible } = action.payload;
      state.header.navigation.visible = visible;
    },
  },
});

export const { setHomeMobileNavigation, handleHomeMobileNavigation } =
  homeSlice.actions;

export default homeSlice.reducer;
