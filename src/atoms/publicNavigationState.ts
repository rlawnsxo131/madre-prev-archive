import { useCallback } from 'react';
import { atom, useRecoilValue, useSetRecoilState } from 'recoil';

interface PublicNavigationState {
  visible: boolean;
}

const publicNavigationState = atom<PublicNavigationState>({
  key: 'publicNavigationState',
  default: {
    visible: false,
  },
});

export function usePublicNavigationValue() {
  return useRecoilValue(publicNavigationState);
}

export function usePublicNavigationActions() {
  const set = useSetRecoilState(publicNavigationState);

  const onClose = useCallback(() => {
    set((prev) => ({
      ...prev,
      visible: false,
    }));
  }, [set]);

  const handleNavigation = useCallback(() => {
    set((prev) => ({
      ...prev,
      visible: !prev.visible,
    }));
  }, [set]);

  return {
    onClose,
    handleNavigation,
  };
}
