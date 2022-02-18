import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import { setHomeMobileNavigation } from '../../store/home';

export default function useHomeHeaderMobileNavigation() {
  const dispatch = useDispatch<AppDispatch>();
  return () => dispatch(setHomeMobileNavigation());
}
