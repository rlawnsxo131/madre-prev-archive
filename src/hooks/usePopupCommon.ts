import { useDispatch } from 'react-redux';
import { AppDispatch } from '../store';
import { setPopupCommon } from '../store/core';

export default function usePopupCommon() {
  const dispatch = useDispatch<AppDispatch>();

  const showPopup = ({
    visible,
    title = '',
    message,
  }: {
    visible: boolean;
    title?: string;
    message: string;
  }) => {
    dispatch(
      setPopupCommon({
        visible,
        title,
        message,
      }),
    );
  };

  const closePopup = () => {
    dispatch(
      setPopupCommon({
        visible: false,
        title: '',
        message: '',
      }),
    );
  };

  return {
    showPopup,
    closePopup,
  };
}
