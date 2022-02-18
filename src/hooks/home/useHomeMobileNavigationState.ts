import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function useHomeMobileNavigationState() {
  return useSelector((state: RootState) => state.home.mobileNavigationState);
}
