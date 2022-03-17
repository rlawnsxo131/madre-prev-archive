import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import home from '../../store/home';

export default function useHomeHeaderActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      handleMobileNavigation() {
        dispatch(home.actions.handleMobileNavigation());
      },
      closeMobileNavigation() {
        dispatch(
          home.actions.setMobileNavigation({
            visible: false,
          }),
        );
      },
    }),
    [dispatch],
  );
}
