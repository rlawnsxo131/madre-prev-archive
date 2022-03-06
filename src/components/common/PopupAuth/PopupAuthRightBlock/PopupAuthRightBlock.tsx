import PopupAuthRightBlockDefault from './PopupAuthRightBlockDefault';
import PopupAuthRightBlockIsError from './PopupAuthRightBlockIsError';

interface PopupAuthRightBlockProps {
  isError: boolean;
  close: () => void;
  resetError: () => void;
}

function PopupAuthRightBlock({
  isError,
  close,
  resetError,
}: PopupAuthRightBlockProps) {
  if (isError) {
    return <PopupAuthRightBlockIsError close={close} resetError={resetError} />;
  }
  return <PopupAuthRightBlockDefault close={close} />;
}

export default PopupAuthRightBlock;
