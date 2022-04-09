import { useMemo } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import user from '../../store/user';

interface SetUserParams {
  access_token: string;
  display_name: string;
}

export default function useUserActions() {
  const dispatch = useDispatch<AppDispatch>();

  return useMemo(
    () => ({
      setUser({ access_token, display_name }: SetUserParams) {
        dispatch(
          user.actions.setUser({
            access_token,
            display_name,
          }),
        );
      },
      setIsPending(isPending: boolean) {
        dispatch(user.actions.setIsPending({ isPending }));
      },
    }),
    [dispatch],
  );
}
