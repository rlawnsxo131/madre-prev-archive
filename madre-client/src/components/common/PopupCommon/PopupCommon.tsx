import usePopupCommonState from '../../../hooks/common/usePopupCommonState';
import usePopupCommonActions from '../../../hooks/common/usePopupCommonActions';
import PopupOkCancel from '../PopupOkCancel';

interface PopupCommonProps {}

function PopupCommon(props: PopupCommonProps) {
  const { close } = usePopupCommonActions();
  const { visible, title, message } = usePopupCommonState();

  return (
    <PopupOkCancel
      visible={visible}
      title={title}
      message={message}
      onConfirm={close}
    />
  );
}

export default PopupCommon;
