import { useDispatch } from 'react-redux';
import { setTheme } from '../../store/theme';

export default function useTheme() {
  const dispatch = useDispatch();

  const handleColorTheme = () => {
    dispatch(setTheme());
  };

  return {
    handleColorTheme,
  };
}
