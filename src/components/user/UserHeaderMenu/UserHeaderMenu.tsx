import { css } from '@emotion/react';
import usePopupAuthActions from '../../../hooks/popupAuth/usePopupAuthActions';
import useUserLoadStatusState from '../../../hooks/user/useUserLoadStatusState';
import useUserMenuButtonActions from '../../../hooks/user/useUserMenuButtonActions';
import useUserMenuState from '../../../hooks/user/useUserMenuState';
import useUserProfileState from '../../../hooks/user/useUserProfileState';
import useUserSignOut from '../../../hooks/user/useUserSignOut';
import UserHeaderMenuAuthButton from './UserHeaderMenuAuthButton';
import UserHeaderMenuIcon from './UserHeaderMenuIcon';
import UserHeaderMenuItems from './UserHeaderMenuItems';

interface UserHeaderMenuProps {}

function UserHeaderMenu(props: UserHeaderMenuProps) {
  const { isPending, isError } = useUserLoadStatusState();
  const menu = useUserMenuState();
  const profile = useUserProfileState();
  const signOut = useUserSignOut();
  const { show } = usePopupAuthActions();
  const { handleNavigation } = useUserMenuButtonActions();

  if (isPending) {
    return <div css={flexCenter}>loading...</div>;
  }

  if (isError) {
    return <div css={flexCenter}>error</div>;
  }

  if (!profile) {
    return (
      <div css={[block, flexCenter]}>
        <UserHeaderMenuAuthButton show={show} />
      </div>
    );
  }

  return (
    <div css={[block, flexCenter]}>
      <UserHeaderMenuIcon
        onClick={handleNavigation}
        photo_url={profile.photo_url}
      />
      <UserHeaderMenuItems
        signOut={signOut}
        visible={menu.visible}
        display_name={profile.display_name}
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
