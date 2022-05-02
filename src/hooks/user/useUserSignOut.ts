import { useLocation, useNavigate, matchRoutes } from 'react-router-dom';
import authApi from '../../store/api/authApi';

const userPersonalPath = '/@:displayName';
const routes = [{ path: userPersonalPath }];

export default function useUserSignOut() {
  const navigate = useNavigate();
  const location = useLocation();
  const routeInfo = matchRoutes(routes, location);
  const [signOut] = authApi.useDeleteMutation();

  return async () => {
    try {
      await signOut(undefined);
      if (routeInfo) {
        const [
          {
            route: { path },
          },
        ] = routeInfo;
        if (path === userPersonalPath) {
          navigate('/', { replace: true });
        }
      }
    } catch (e) {}
  };
}
