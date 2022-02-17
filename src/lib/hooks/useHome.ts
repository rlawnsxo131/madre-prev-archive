import { useDispatch } from 'react-redux';
import { setMobileNavigation } from '../../store/home';

export default function useHome() {
  const dispatch = useDispatch();

  const handleMobileNavigation = () => {
    dispatch(setMobileNavigation());
  };

  return {
    handleMobileNavigation,
  };
}
