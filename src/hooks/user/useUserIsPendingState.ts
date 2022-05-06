import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function useUserIsPendingState() {
  return useSelector((state: RootState) => state.user.isPending);
}
