import { useEffect, useMemo, useTransition } from 'react';
import { useDispatch } from 'react-redux';
import { MADRE_USER_PROFILE } from '../../constants';
import { Storage } from '../../lib/storage';
import { AppDispatch } from '../../store';
import authApi from '../../store/api/authApi';
import user from '../../store/user';

export default function useUserLoadEffect() {
  const dispatch = useDispatch<AppDispatch>();
  const [isPending, startTransition] = useTransition();
  const { isFetching } = authApi.useGetQuery(undefined);

  const isPendingVisible = useMemo(() => {
    return isFetching || isPending;
  }, [isFetching, isPending]);

  useEffect(() => {
    startTransition(() => {
      const profile = Storage.getItem(MADRE_USER_PROFILE);
      if (!profile) return;
      dispatch(
        user.actions.setUser({
          profile,
        }),
      );
    });
  }, [dispatch]);

  useEffect(() => {
    dispatch(
      user.actions.setLoadUserStatusIsPending({
        isPending: isPendingVisible,
      }),
    );
  }, [dispatch, isPendingVisible]);
}
