import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../redux';
import { setTheme } from '../../redux/theme';

export default function useSetTheme() {
  const dispatch = useDispatch<AppDispatch>();
  return () => dispatch(setTheme());
}
