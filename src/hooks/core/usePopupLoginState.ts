import { useSelector } from 'react-redux';
import { RootState } from '../../redux';

export default function usePopupLoginState() {
  return useSelector((state: RootState) => state.core.popupLogin);
}
