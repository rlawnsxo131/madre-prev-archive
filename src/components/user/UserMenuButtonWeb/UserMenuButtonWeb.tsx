import { css } from '@emotion/react';
import useUserState from '../../../hooks/user/useUserState';
import { themePalette } from '../../../styles';
import UserMenuButtonProfileIcon from './UserMenuButtonProfileIcon';
import UserMenuButtonAuth from './UserMenuButtonWebAuth';
import UserMenuButtonWebNavigation from './UserMenuButtonWebNavigation';

interface UserMenuButtonWebProps {}

function UserMenuButtonWeb(props: UserMenuButtonWebProps) {
  const { isPending, userTokenProfile } = useUserState();

  if (isPending) {
    return (
      <div style={{ display: 'flex', alignItems: 'center' }}>loading...</div>
    );
  }

  if (!userTokenProfile) {
    return (
      <div css={block}>
        <UserMenuButtonAuth />
      </div>
    );
  }

  return (
    <div css={block}>
      <UserMenuButtonProfileIcon />
      <UserMenuButtonWebNavigation />
    </div>
  );
}

const block = css`
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.4rem;
  cursor: pointer;

  img {
    width: 1.9rem;
    height: 1.9rem;
    border-radius: 100%;
  }

  svg {
    fill: ${themePalette.fill1};
    &:nth-of-type(1) {
      width: 1.9rem;
      height: 1.9rem;
    }
    &:nth-of-type(2) {
      width: 0.5rem;
      height: 0.5rem;
    }
  }
`;

export default UserMenuButtonWeb;
