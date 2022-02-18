import { useDispatch } from 'react-redux';
import { AppDispatch } from '../store';
import { setMobileNavigation } from '../store/home';

export default function useHome() {
  const dispatch = useDispatch<AppDispatch>();

  const handleMobileNavigation = () => {
    dispatch(setMobileNavigation());
  };

  return {
    handleMobileNavigation,
  };
}
