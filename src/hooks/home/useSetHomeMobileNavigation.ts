import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../redux';
import { setHomeMobileNavigation } from '../../redux/home';

export default function useHomeHeaderMobileNavigation() {
  const dispatch = useDispatch<AppDispatch>();
  return () => dispatch(setHomeMobileNavigation());
}
