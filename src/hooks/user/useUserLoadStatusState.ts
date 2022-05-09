import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function useUserLoadStatusState() {
  return useSelector((state: RootState) => state.user.loadUserStatus);
}
