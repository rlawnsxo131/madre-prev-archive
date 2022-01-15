import { css } from '@emotion/react';
import GoogleAuthButton from '../../auth/GoogleAuthButton';

interface UserPersonalMenuProps {}

function UserPersonalMenu(props: UserPersonalMenuProps) {
  return (
    <div css={block}>
      <GoogleAuthButton />
    </div>
  );
}

const block = css`
  position: relative;
`;

export default UserPersonalMenu;
