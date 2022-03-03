import Button from '../Button';
import { ButtonSize } from '../Button/Button.styles';
import { GoogleIcon } from '../../../image/icons';
import useButtonGoogleSignin from './hooks/useButtonGoogleSignin';

interface ButtonGoogleSigninProps {
  size?: ButtonSize;
}

function ButtonGoogleSignin({ size = 'medium' }: ButtonGoogleSigninProps) {
  const { buttonRef } = useButtonGoogleSignin();

  return (
    <Button buttonRef={buttonRef} size={size} icon={<GoogleIcon />} outline>
      Signin with Google
    </Button>
  );
}

export default ButtonGoogleSignin;
