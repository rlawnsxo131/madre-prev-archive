import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { UserTokenProfile } from '../../@types/domain/auth.type';
import { AppDispatch } from '../../store';
import user from '../../store/user';

export default function useUserActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      setUser(userProfile: UserTokenProfile) {
        dispatch(user.actions.setUser({ userProfile }));
      },
      setIsPending(isPending: boolean) {
        dispatch(user.actions.setIsPending({ isPending }));
      },
    }),
    [dispatch],
  );
}
