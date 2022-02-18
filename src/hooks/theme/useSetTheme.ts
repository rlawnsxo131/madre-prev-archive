import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import { setTheme } from '../../store/theme';

export default function useSetTheme() {
  const dispatch = useDispatch<AppDispatch>();
  return () => dispatch(setTheme());
}
