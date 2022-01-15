import { GoogleIcon } from '../../../image/icons';
import Button from '../../common/Button';
import useGoogleAuthButton from './hooks/useGoogleAuthButton';

interface GoogleAuthButtonProps {}

function GoogleAuthButton(props: GoogleAuthButtonProps) {
  const { buttonRef } = useGoogleAuthButton();
  return (
    <Button buttonRef={buttonRef} size="medium" icon={<GoogleIcon />} outline>
      Sign in with Google
    </Button>
  );
}

export default GoogleAuthButton;
