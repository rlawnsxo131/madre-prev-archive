import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import user from '../../store/user';

export default function useUserMenuButtonActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      handleNavigation() {
        dispatch(user.actions.handleNavigation());
      },
    }),
    [dispatch],
  );
}
