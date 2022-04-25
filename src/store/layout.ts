import { createSlice } from '@reduxjs/toolkit';

interface LayoutState {
  header: {
    navigation: {
      visible: boolean;
    };
  };
}

const initialState: LayoutState = {
  header: {
    navigation: {
      visible: false,
    },
  },
};

const layout = createSlice({
  name: 'layout',
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

export default layout;
