import { useCallback } from 'react';
import { atom, useRecoilValue, useSetRecoilState } from 'recoil';

interface AppNavigationState {
  visible: boolean;
}

const appNavigationState = atom<AppNavigationState>({
  key: 'appNavigationState',
  default: {
    visible: false,
  },
});

export function useAppNavigationValue() {
  return useRecoilValue(appNavigationState);
}

export function useAppNavigationActions() {
  const set = useSetRecoilState(appNavigationState);

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
