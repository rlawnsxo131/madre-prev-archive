import { useCallback } from 'react';
import { atom, selector, useRecoilValue, useSetRecoilState } from 'recoil';
import { darkmodeStateSelector, DarkmodeThemeType } from './darkmodeState';

interface NavigationState {
  visible: boolean;
}

const navigationState = atom<NavigationState>({
  key: 'navigationState',
  default: {
    visible: false,
  },
});

interface HomeState {
  visible: boolean;
  theme: DarkmodeThemeType;
}

const homeStateSelector = selector<HomeState>({
  key: 'homeStateSelector',
  get: ({ get }) => {
    const { theme } = get(darkmodeStateSelector);
    const { visible } = get(navigationState);
    return {
      theme,
      visible,
    };
  },
});

export function useHomeState() {
  const { theme, visible } = useRecoilValue(homeStateSelector);
  const set = useSetRecoilState(navigationState);

  const handleVisible = useCallback(() => {
    set((prev) => ({
      ...prev,
      visible: !prev.visible,
    }));
  }, [set]);

  return {
    theme,
    visible,
    handleVisible,
  };
}
