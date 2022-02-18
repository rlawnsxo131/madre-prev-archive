import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import { setPopupCommon } from '../../store/core';

export default function useClosePopupCommon() {
  const dispatch = useDispatch<AppDispatch>();
  return () =>
    dispatch(
      setPopupCommon({
        visible: false,
        title: '',
        message: '',
      }),
    );
}
