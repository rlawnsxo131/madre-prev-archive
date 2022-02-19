import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../redux';
import { setHomeMobileNavigation } from '../../redux/home';

export default function useHomeHeaderActions() {
  const dispatch = useDispatch<AppDispatch>();

  const handleMobileNavigation = () => dispatch(setHomeMobileNavigation());

  return {
    handleMobileNavigation,
  };
}
