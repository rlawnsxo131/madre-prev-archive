import useUserMenuButtonActions from '../../../hooks/user/useUserMenuButtonActions';
import { UserIcon } from '../../../image/icons';
import { googlePhotoUrlSizeChange } from '../../../lib/utils';

interface UserMenuButtonProfileIconProps {
  photo_url?: string;
}

function UserMenuButtonProfileIcon({
  photo_url,
}: UserMenuButtonProfileIconProps) {
  const { handleNavigation } = useUserMenuButtonActions();

  if (!photo_url) {
    return <UserIcon onClick={handleNavigation} />;
  }

  return (
    <img src={googlePhotoUrlSizeChange(photo_url)} onClick={handleNavigation} />
  );
}

export default UserMenuButtonProfileIcon;
