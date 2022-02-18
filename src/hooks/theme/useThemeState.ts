import { useSelector } from 'react-redux';
import { RootState } from '../../redux';

export default function useThemeState() {
  return useSelector((state: RootState) => state.theme);
}
