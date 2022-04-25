import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import layout from '../../store/layout';

export default function useLayoutHeaderActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      handleMobileNavigation() {
        dispatch(layout.actions.handleMobileNavigation());
      },
      closeMobileNavigation() {
        dispatch(layout.actions.closeMobileNavigation());
      },
    }),
    [dispatch],
  );
}
