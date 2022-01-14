import { GoogleIcon } from '../../../image/icons';
import Button from '../../common/Button';
import useAuthGoogleButton from './hooks/useAuthGoogleButton';

interface AuthGoogleButtonProps {}

function AuthGoogleButton(props: AuthGoogleButtonProps) {
  const { buttonRef } = useAuthGoogleButton();
  return (
    <Button buttonRef={buttonRef} size="medium" icon={<GoogleIcon />} outline>
      Sign in with Google
    </Button>
  );
}

export default AuthGoogleButton;
