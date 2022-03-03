import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function useScreenSignupState() {
  return useSelector((state: RootState) => state.core.screenSignup);
}
