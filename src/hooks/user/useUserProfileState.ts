import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function useUserProfileState() {
  return useSelector((state: RootState) => state.user.profile);
}
