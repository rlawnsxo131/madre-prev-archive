import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export default function usePopupCommonState() {
  return useSelector((state: RootState) => state.core.popupCommon);
}
