import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function useLoadingState() {
  return useSelector((state: RootState) => state.core.loading);
}
