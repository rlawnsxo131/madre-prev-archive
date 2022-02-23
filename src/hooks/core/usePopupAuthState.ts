import { useSelector } from 'react-redux';
import { RootState } from '../../redux';

export default function usePopupAuthState() {
  return useSelector((state: RootState) => state.core.popupAuth);
}
