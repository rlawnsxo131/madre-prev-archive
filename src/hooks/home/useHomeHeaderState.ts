import { useSelector } from 'react-redux';
import { RootState } from '../../redux';

export default function useHomeHeaderState() {
  return useSelector((state: RootState) => state.home.header);
}
