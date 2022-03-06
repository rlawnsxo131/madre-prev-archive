import { CloseIcon } from '../../../../image/icons';
import { InspirationImage } from '../../../../image/images';
import ButtonGoogleSignin from '../../ButtonGoogleSignin';
import PopupAuthStyles from '../PopupAuth.styles';

interface PopupAuthRightBlockDefaultProps {
  close: () => void;
}

function PopupAuthRightBlockDefault({
  close,
}: PopupAuthRightBlockDefaultProps) {
  return (
    <div css={PopupAuthStyles.rightBlock}>
      <div css={PopupAuthStyles.rightBlockHeader}>
        <CloseIcon className="popup-auth-close-icon" onClick={close} />
        <h1>Welcome To Madre</h1>
      </div>
      <div css={PopupAuthStyles.rightBlockBody}>
        <InspirationImage />
      </div>
      <ButtonGoogleSignin />
    </div>
  );
}

export default PopupAuthRightBlockDefault;
