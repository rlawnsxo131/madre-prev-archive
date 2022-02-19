import { GoogleIcon } from '../../../image/icons';
import Button from '../../common/Button';
import { ButtonSize } from '../../common/Button/Button.styles';
import useGoogleAuthButton from './hooks/useGoogleAuthButton';

interface GoogleAuthButtonProps {
  size?: ButtonSize;
}

function GoogleAuthButton({ size = 'medium' }: GoogleAuthButtonProps) {
  const { buttonRef } = useGoogleAuthButton();
  return (
    <Button buttonRef={buttonRef} size={size} icon={<GoogleIcon />} outline>
      Login with Google
    </Button>
  );
}

export default GoogleAuthButton;
