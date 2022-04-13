import { useEffect, useMemo, useTransition } from 'react';
import { useSelector } from 'react-redux';
import { useDispatch } from 'react-redux';
import { MADRE_USER } from '../../constants';
import { Storage } from '../../lib/storage';
import { AppDispatch, RootState } from '../../store';
import authApi from '../../store/api/authApi';
import user from '../../store/user';

export default function useUserLoadEffect() {
  const dispatch = useDispatch<AppDispatch>();
  const prevUser = useSelector((state: RootState) => state.user);
  const [isPending, startTransition] = useTransition();
  const { data } = authApi.useGetQuery(null);

  useEffect(() => {
    startTransition(() => {
      const userData = Storage.getItem(MADRE_USER);
      if (!userData) return;
      dispatch(
        user.actions.setUser({
          userProfile: userData,
        }),
      );
    });
  }, [dispatch]);

  useEffect(() => {
    console.log(data);
  }, [data]);

  useEffect(() => {
    dispatch(user.actions.setIsPending({ isPending }));
  }, [dispatch]);
}
