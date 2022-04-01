import Button from '../Button';
import { ButtonSize } from '../Button/Button.styles';
import { GoogleIcon } from '../../../image/icons';
import useButtonGoogleSignIn from './hooks/useButtonGoogleSignIn';

interface ButtonGoogleSignInProps {
  size?: ButtonSize;
}

function ButtonGoogleSignIn({ size = 'medium' }: ButtonGoogleSignInProps) {
  const { buttonRef } = useButtonGoogleSignIn();

  return (
    <Button buttonRef={buttonRef} size={size} icon={<GoogleIcon />} outline>
      SignIn with Google
    </Button>
  );
}

export default ButtonGoogleSignIn;
