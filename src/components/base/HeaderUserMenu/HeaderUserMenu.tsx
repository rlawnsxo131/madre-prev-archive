import { css } from '@emotion/react';
import usePopupAuthActions from '../../../hooks/popupAuth/usePopupAuthActions';
import useUserLoadStatusState from '../../../hooks/user/useUserLoadStatusState';
import useUserMenuButtonActions from '../../../hooks/user/useUserMenuButtonActions';
import useUserMenuState from '../../../hooks/user/useUserMenuState';
import useUserProfileState from '../../../hooks/user/useUserProfileState';
import useUserSignOut from '../../../hooks/user/useUserSignOut';
import HeaderUserMenuAuthButton from './HeaderUserMenuAuthButton';
import HeaderUserMenuIcon from './HeaderUserMenuIcon';
import HeaderUserMenuItems from './HeaderUserMenuItems';

interface HeaderUserMenuProps {}

function HeaderUserMenu(props: HeaderUserMenuProps) {
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
        <HeaderUserMenuAuthButton show={show} />
      </div>
    );
  }

  return (
    <div css={[block, flexCenter]}>
      <HeaderUserMenuIcon
        onClick={handleNavigation}
        photo_url={profile.photo_url}
      />
      <HeaderUserMenuItems
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

export default HeaderUserMenu;
