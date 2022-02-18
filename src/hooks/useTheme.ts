import { useDispatch } from 'react-redux';
import { AppDispatch } from '../store';
import { setTheme } from '../store/theme';

export default function useTheme() {
  const dispatch = useDispatch<AppDispatch>();

  const handleColorTheme = () => {
    dispatch(setTheme());
  };

  return {
    handleColorTheme,
  };
}
