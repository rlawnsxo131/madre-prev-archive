import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import theme from '../../store/theme';

export default function useThemeActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      handleTheme() {
        dispatch(theme.actions.handleTheme());
      },
    }),
    [dispatch],
  );
}
