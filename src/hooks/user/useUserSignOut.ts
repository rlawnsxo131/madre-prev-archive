import { useNavigate } from 'react-router-dom';
import authApi from '../../store/api/authApi';
import useIsUserPath from '../useIsUserPath';

export default function useUserSignOut() {
  const navigate = useNavigate();
  const [signOut] = authApi.useDeleteMutation();
  const isUserPath = useIsUserPath();

  return async () => {
    try {
      await signOut(undefined);
      if (isUserPath) {
        navigate('/', { replace: true });
      }
    } catch (e) {}
  };
}
