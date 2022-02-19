import { useCallback } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../redux';
import { setTheme } from '../../redux/theme';

export default function useThemeActions() {
  const dispatch = useDispatch<AppDispatch>();

  const handleTheme = useCallback(() => {
    dispatch(setTheme());
  }, [dispatch]);

  return {
    handleTheme,
  };
}
