import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../redux';
import { setPopupCommon } from '../../redux/core';

export default function usePopupCommonClose() {
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
