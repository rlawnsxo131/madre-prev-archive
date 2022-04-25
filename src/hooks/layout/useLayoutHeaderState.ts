import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function useHeaderState() {
  return useSelector((state: RootState) => state.layout.header);
}
