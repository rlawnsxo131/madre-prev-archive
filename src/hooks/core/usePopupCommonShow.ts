import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../redux';
import { setPopupCommon } from '../../redux/core';

interface UsePopupCommonShowParams {
  title?: string;
  message: string;
}

export default function usePopupCommonShow({
  title = '',
  message,
}: UsePopupCommonShowParams) {
  const dispatch = useDispatch<AppDispatch>();
  return () => {
    dispatch(
      setPopupCommon({
        visible: true,
        title,
        message,
      }),
    );
  };
}
