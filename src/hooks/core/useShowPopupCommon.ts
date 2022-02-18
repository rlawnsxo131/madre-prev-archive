import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../store';
import { setPopupCommon } from '../../store/core';

interface UseShowPopupCommonParams {
  title?: string;
  message: string;
}

export default function useShowPopupCommon({
  title = '',
  message,
}: UseShowPopupCommonParams) {
  const dispatch = useDispatch<AppDispatch>();
  return () =>
    dispatch(
      setPopupCommon({
        visible: true,
        title,
        message,
      }),
    );
}
