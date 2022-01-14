import { useCallback } from 'react';
import { atom, useRecoilValue, useSetRecoilState } from 'recoil';

interface HomeState {
  navigation: {
    visible: boolean;
  };
}

const homeState = atom<HomeState>({
  key: 'homeState',
  default: {
    navigation: {
      visible: false,
    },
  },
});

export function useHomeValue() {
  return useRecoilValue(homeState);
}

export function useHomeActions() {
  const set = useSetRecoilState(homeState);

  const handleNavigation = useCallback(() => {
    set((prev) => ({
      ...prev,
      navigation: {
        ...prev.navigation,
        visible: !prev.navigation.visible,
      },
    }));
  }, [set]);

  return {
    handleNavigation,
  };
}
