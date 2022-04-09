import { useEffect, useTransition } from 'react';
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

  const { isFetching, isError, data } = authApi.useGetQuery({});

  console.log(data);

  useEffect(() => {
    startTransition(() => {
      const userData = Storage.getItem(MADRE_USER);
      if (!userData) return;
      dispatch(
        user.actions.setUser({
          access_token: userData.access_token,
          display_name: userData.display_name,
        }),
      );
    });
  }, [dispatch]);

  useEffect(() => {
    dispatch(user.actions.setIsPending({ isPending }));
  }, [dispatch, isPending]);
}
