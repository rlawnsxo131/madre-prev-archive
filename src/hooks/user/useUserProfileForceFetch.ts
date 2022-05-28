import { useCallback } from 'react';
import { useDispatch } from 'react-redux';
import authApi from '../../store/api/authApi';

export default function useUserProfileForceFetch() {
  const dispatch = useDispatch();

  return useCallback(() => {
    dispatch(
      authApi.endpoints.get.initiate(undefined, {
        subscribe: false,
        forceRefetch: true,
      }),
    );
  }, [dispatch]);
}
