import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import { setTheme } from '../../store/theme';

export default function useThemeActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      handleTheme() {
        dispatch(setTheme());
      },
    }),
    [dispatch],
  );
}
