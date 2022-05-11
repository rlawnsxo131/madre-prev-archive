import { css } from '@emotion/react';
import { UserIcon } from '../../../image/icons';
import { basicStyles } from '../../../styles';

interface FooterMobileUserMenuProps {}

function FooterMobileUserMenu(props: FooterMobileUserMenuProps) {
  // const navigate = useNavigate();
  // const profile = useUserProfileState();
  // const { show } = usePopupAuthActions();
  // const onClickMobileUserMenu = () => {
  //   if (!profile || !profile?.username) {
  //     show();
  //     return;
  //   }
  //   navigate(`/@${profile.username}`);
  // };
  return (
    <button css={[basicStyles.button, button]}>
      <UserIcon />
    </button>
  );
}

const button = css`
  display: flex;
  justify-content: center;
  align-items: center;
`;

export default FooterMobileUserMenu;
