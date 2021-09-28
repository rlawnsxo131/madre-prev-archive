import { useCallback } from 'react';
import { atom, selector, useRecoilValue, useSetRecoilState } from 'recoil';
import { darkmodeState, DarkmodeTheme } from './darkmodeState';

interface HomeNavigationState {
  visible: boolean;
}

const homeNavigationState = atom<HomeNavigationState>({
  key: 'navigationState',
  default: {
    visible: false,
  },
});

interface HomeState {
  visible: boolean;
  theme: DarkmodeTheme;
}

const homeStateSelector = selector<HomeState>({
  key: 'homeStateSelector',
  get: ({ get }) => {
    const { visible } = get(homeNavigationState);
    const { theme } = get(darkmodeState);
    return {
      visible,
      theme,
    };
  },
});

export function useHomeValue() {
  return useRecoilValue(homeStateSelector);
}

export function useHomeActions() {
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
