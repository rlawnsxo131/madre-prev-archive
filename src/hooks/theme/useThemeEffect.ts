import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from '../../store';
import theme from '../../store/theme';

export default function useThemeEffect() {
  const dispatch = useDispatch<AppDispatch>();
  const { theme: currentTheme } = useSelector(
    (state: RootState) => state.theme,
  );

  useEffect(() => {
    const systemPrefersDark = window.matchMedia(
      '(prefers-color-scheme: dark)',
    ).matches;
    if (systemPrefersDark) {
      dispatch(theme.actions.setTheme('dark'));
    }
  }, [dispatch]);

  useEffect(() => {
    document.body.dataset.theme = currentTheme;
  }, [currentTheme]);
}
