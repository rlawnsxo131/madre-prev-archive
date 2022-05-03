import { useMemo } from 'react';
import { useLocation, matchRoutes } from 'react-router-dom';
import { appRoutes } from '../constants';

export default function useMatchedRoute() {
  const location = useLocation();
  const routeInfo = matchRoutes(appRoutes, location);

  return useMemo(() => {
    if (routeInfo) {
      const [
        {
          route: { path },
        },
      ] = routeInfo;
      return path;
    }
    return null;
  }, [routeInfo]);
}
