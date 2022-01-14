import { useCallback } from 'react';
import { atom, useRecoilValue, useSetRecoilState } from 'recoil';

interface HomeNavigationState {
  visible: boolean;
}

const homeNavigationState = atom<HomeNavigationState>({
  key: 'homeNavigationState',
  default: {
    visible: false,
  },
});

export function useHomeNavigationValue() {
  return useRecoilValue(homeNavigationState);
}

export function useHomeNavigationActions() {
  const set = useSetRecoilState(homeNavigationState);

  const handleNavigation = useCallback(() => {
    set((prev) => ({
      ...prev,
      visible: !prev.visible,
    }));
  }, [set]);

  return {
    handleNavigation,
  };
}
