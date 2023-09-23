import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function useUserMenuState() {
  return useSelector((state: RootState) => state.user.menu);
}
