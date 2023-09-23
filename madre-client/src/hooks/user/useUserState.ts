import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function useUserState() {
  return useSelector((state: RootState) => state.user);
}
