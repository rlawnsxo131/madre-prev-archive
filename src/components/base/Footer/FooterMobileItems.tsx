import { css } from '@emotion/react';
import { useNavigate } from 'react-router-dom';
import { userPath } from '../../../constants';
import usePopupAuthActions from '../../../hooks/popupAuth/usePopupAuthActions';
import useUserProfileState from '../../../hooks/user/useUserProfileState';
import { HomeIcon, MenuIcon, UserIcon } from '../../../image/icons';
import MadreButtonLink from '../../common/MadreButtonLink';
import MadreImageLink from '../../common/MadreImageLink';
import FooterMobileNotification from './FooterMobileNotification';

interface FooterMobileItemsProps {}

function FooterMobileItems(props: FooterMobileItemsProps) {
  const navigate = useNavigate();
  const profile = useUserProfileState();
  const { show } = usePopupAuthActions();
  const onClickMobileUserMenu = () => {
    if (!profile || !profile?.username) {
      show();
      return;
    }
    navigate(`/@${profile.username}`);
  };

  return (
    <div css={block}>
      <MadreImageLink to="/">
        <HomeIcon />
      </MadreImageLink>
      <MadreButtonLink onClick={onClickMobileUserMenu} matchPath={userPath}>
        <UserIcon />
      </MadreButtonLink>
      <MadreImageLink to="/notifications">
        <FooterMobileNotification />
      </MadreImageLink>
      <MadreImageLink to="/m/all-menu">
        <MenuIcon />
      </MadreImageLink>
    </div>
  );
}

const block = css`
  width: 93%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 3rem;
`;

export default FooterMobileItems;
