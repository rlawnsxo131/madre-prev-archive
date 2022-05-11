import { useMemo } from 'react';
import { userPath } from '../constants';
import useMatchedRoute from './useMatchedRoute';

export default function useIsUserPath() {
  const matchPath = useMatchedRoute();

  return useMemo(() => {
    return matchPath?.startsWith(userPath) ?? false;
  }, [matchPath]);
}
