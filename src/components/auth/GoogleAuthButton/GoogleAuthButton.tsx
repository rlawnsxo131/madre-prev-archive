import { css } from '@emotion/react';
import { GoogleIcon } from '../../../image/icons';
import Button from '../../common/Button';
import useGoogleAuthButton from './hooks/useGoogleAuthButton';

interface GoogleAuthButtonProps {}

function GoogleAuthButton(props: GoogleAuthButtonProps) {
  const { buttonRef } = useGoogleAuthButton();
  return (
    <div css={block}>
      <Button buttonRef={buttonRef} size="medium" icon={<GoogleIcon />} outline>
        Sign in with Google
      </Button>
    </div>
  );
}

const block = css`
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0 0.5rem;
`;

export default GoogleAuthButton;
