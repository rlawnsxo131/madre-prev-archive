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

const home = createSlice({
  name: 'home',
  initialState,
  reducers: {
    handleMobileNavigation(state) {
      const { visible } = state.header.navigation;
      state.header.navigation.visible = !visible;
    },
    closeMobileNavigation(state) {
      state.header.navigation.visible = false;
    },
  },
});

export default home;
