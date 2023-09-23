import { CloseIcon } from '../../../image/icons';
import { InspirationImage } from '../../../image/images';
import ButtonGoogleSignIn from '../ButtonGoogleSignIn';
import PopupAuthStyles from './PopupAuth.styles';

interface PopupAuthRightBlockWelcomeSignInProps {
  close: () => void;
}

function PopupAuthRightBlockWelcomeSignIn({
  close,
}: PopupAuthRightBlockWelcomeSignInProps) {
  return (
    <div css={PopupAuthStyles.rightBlock}>
      <div css={PopupAuthStyles.rightBlockHeader}>
        <CloseIcon className="popup-auth-close-icon" onClick={close} />
        <h1>Madre 에 오신걸 환영해요!</h1>
      </div>
      <div css={PopupAuthStyles.rightBlockBody}>
        <InspirationImage />
      </div>
      <ButtonGoogleSignIn />
    </div>
  );
}

export default PopupAuthRightBlockWelcomeSignIn;
