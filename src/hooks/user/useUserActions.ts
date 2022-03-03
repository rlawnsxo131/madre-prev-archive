import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';

export default function useUserActions() {
  const dispatch = useDispatch<AppDispatch>();

  return {};
}
