import { useCallback } from 'react';
import { atom, useRecoilValue, useSetRecoilState } from 'recoil';

interface HomeHeaderMobileNavigationState {
  visible: boolean;
}

const mobilePublicNavigationState = atom<HomeHeaderMobileNavigationState>({
  key: 'mobilePublicNavigationState',
  default: {
    visible: false,
  },
});

export function useHomeHeaderMobileNavigationValue() {
  return useRecoilValue(mobilePublicNavigationState);
}

export function useHomeHeaderMobileNavigationActions() {
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
