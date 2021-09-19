import { useCallback } from 'react';
import { atom, selector, useRecoilValue, useSetRecoilState } from 'recoil';
import { useDarkmodeState } from './darkmodeState';

interface NavigationState {
  visible: boolean;
}

const navigationState = atom<NavigationState>({
  key: 'navigationState',
  default: {
    visible: false,
  },
});

const homeStateSelector = selector<NavigationState>({
  key: 'homeStateSelector',
  get: ({ get }) => {
    const { visible } = get(navigationState);
    return {
      visible,
    };
  },
});

export function useHomeState() {
  const { theme } = useDarkmodeState();
  const { visible } = useRecoilValue(homeStateSelector);
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
