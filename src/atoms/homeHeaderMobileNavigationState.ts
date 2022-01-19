import { useCallback } from 'react';
import { atom, useRecoilValue, useSetRecoilState } from 'recoil';

interface HomeHeaderMobileNavigationState {
  visible: boolean;
}

const homeHeaderMobileNavigationState = atom<HomeHeaderMobileNavigationState>({
  key: 'mobilePublicNavigationState',
  default: {
    visible: false,
  },
});

export function useHomeHeaderMobileNavigationValue() {
  return useRecoilValue(homeHeaderMobileNavigationState);
}

export function useHomeHeaderMobileNavigationActions() {
  const set = useSetRecoilState(homeHeaderMobileNavigationState);

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
