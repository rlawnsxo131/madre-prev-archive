import { useCallback } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../redux';
import { setHomeMobileNavigation } from '../../redux/home';

export default function useHomeHeaderActions() {
  const dispatch = useDispatch<AppDispatch>();

  const handleMobileNavigation = useCallback(() => {
    dispatch(setHomeMobileNavigation());
  }, [dispatch]);

  return {
    handleMobileNavigation,
  };
}
