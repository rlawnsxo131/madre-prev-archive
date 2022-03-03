import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function usePopupAuthState() {
  return useSelector((state: RootState) => state.core.popupAuth);
}
