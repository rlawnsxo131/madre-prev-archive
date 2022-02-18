import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function useThemeState() {
  return useSelector((state: RootState) => state.theme);
}
