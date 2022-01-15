import { useCallback } from 'react';
import { atom, useRecoilValue, useSetRecoilState } from 'recoil';

interface MobilePublicNavigationState {
  visible: boolean;
}

const mobilePublicNavigationState = atom<MobilePublicNavigationState>({
  key: 'mobilePublicNavigationState',
  default: {
    visible: false,
  },
});

export function useMobilePublicNavigationValue() {
  return useRecoilValue(mobilePublicNavigationState);
}

export function useMobilePublicNavigationActions() {
  const set = useSetRecoilState(mobilePublicNavigationState);

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
