import { css } from '@emotion/react';
import usePopupAuthActions from '../../../hooks/popupAuth/usePopupAuthActions';
import useUserMenuButtonActions from '../../../hooks/user/useUserMenuButtonActions';
import useUserSignOut from '../../../hooks/user/useUserSignOut';
import useUserState from '../../../hooks/user/useUserState';
import UserHeaderMenuAuthButton from './UserHeaderMenuAuthButton';
import UserHeaderMenuIcon from './UserHeaderMenuIcon';
import UserHeaderMenuItems from './UserHeaderMenuItems';

interface UserHeaderMenuProps {}

function UserHeaderMenu(props: UserHeaderMenuProps) {
  const { isPending, userTokenProfile, menu } = useUserState();
  const signOut = useUserSignOut();
  const { show } = usePopupAuthActions();
  const { handleNavigation } = useUserMenuButtonActions();

  if (isPending) {
    return <div css={flexCenter}>loading...</div>;
  }

  if (!userTokenProfile) {
    return (
      <div css={[block, flexCenter]}>
        <UserHeaderMenuAuthButton show={show} />
      </div>
    );
  }

  const { photo_url, display_name } = userTokenProfile;

  return (
    <div css={[block, flexCenter]}>
      <UserHeaderMenuIcon onClick={handleNavigation} photo_url={photo_url} />
      <UserHeaderMenuItems
        signOut={signOut}
        visible={menu.visible}
        display_name={display_name}
      />
    </div>
  );
}

const flexCenter = css`
  display: flex;
  justify-content: center;
  align-items: center;
`;

const block = css`
  position: relative;
  gap: 0.4rem;
  cursor: pointer;
`;

export default UserHeaderMenu;
