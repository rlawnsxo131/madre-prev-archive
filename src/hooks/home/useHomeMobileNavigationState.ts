import { useSelector } from 'react-redux';
import { RootState } from '../../redux';

export default function useHomeMobileNavigationState() {
  return useSelector((state: RootState) => state.home.mobileNavigationState);
}
