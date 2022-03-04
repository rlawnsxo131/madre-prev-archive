import { useCallback } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';

export default function useUserActions() {
  const dispatch = useDispatch<AppDispatch>();

  const setUser = useCallback(() => {}, [dispatch]);

  return {
    setUser,
  };
}
