import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import {
  handleHomeMobileNavigation,
  setHomeMobileNavigation,
} from '../../store/home';

export default function useHomeHeaderActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      handleMobileNavigation() {
        dispatch(handleHomeMobileNavigation());
      },
      closeMobileNavigation() {
        dispatch(
          setHomeMobileNavigation({
            visible: false,
          }),
        );
      },
    }),
    [dispatch],
  );
}
