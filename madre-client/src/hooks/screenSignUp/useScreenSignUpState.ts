import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function useScreenSignUpState() {
  return useSelector((state: RootState) => state.screenSignUp);
}
