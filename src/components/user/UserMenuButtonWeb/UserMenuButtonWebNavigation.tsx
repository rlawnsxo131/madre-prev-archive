import { css } from '@emotion/react';
import { NavLink } from 'react-router-dom';
import useUserSignOut from '../../../hooks/user/useUserSignOut';
import useUserState from '../../../hooks/user/useUserState';
import useTransitionTimeoutEffect from '../../../hooks/useTransitionTimeoutEffect';
import {
  standardColor,
  themePalette,
  transitions,
  zIndexes,
} from '../../../styles';

interface UserMenuButtonWebNavigationProps {}

function UserMenuButtonWebNavigation(props: UserMenuButtonWebNavigationProps) {
  const signOut = useUserSignOut();
  const { menu, userTokenProfile } = useUserState();
  const closed = useTransitionTimeoutEffect({
    visible: menu.visible,
  });

  if (!menu.visible && closed) return null;
  if (!userTokenProfile) return null;

  return (
    <div css={block(menu.visible)}>
      <ul css={ul}>
        <li>
          <NavLink
            css={link}
            to={`/@${userTokenProfile.display_name}`}
            // style={({isActive}) => ({
            //   color: isActive ? standardColor.navigation.active : themePalette.text1
            // })}
          >
            마이 페이지
          </NavLink>
        </li>
        <li>
          <button css={button} onClick={signOut}>
            로그아웃
          </button>
        </li>
      </ul>
    </div>
  );
}

const block = (visible: boolean) => css`
  position: absolute;
  top: 3.25rem;
  right: 0;
  width: 8rem;
  height: auto;
  padding: 0.25rem 0.5rem;
  display: flex;
  flex-direction: column;
  z-index: ${zIndexes.dropdownItem};
  border-radius: 0.25rem;
  background: ${themePalette.bg_element1};
  box-shadow: ${themePalette.shadow1};
  transform-origin: top;
  ${visible
    ? css`
        animation: ${transitions.scaleUp} 0.25s forwards ease-in-out;
      `
    : css`
        animation: ${transitions.scaleDown} 0.25s forwards ease-in-out;
      `};
`;

const ul = css`
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
`;

const link = css`
  display: flex;
  flex-flow: row wrap;
  align-items: center;
  font-size: 0.9rem;
  font-weight: bold;
  padding: 0.5rem 0.25rem 0.5rem 0.25rem;
  &.active {
    color: ${standardColor.navigation.active};
  }
`;

const button = css`
  background: none;
  outline: none;
  border: none;
  box-sizing: border-box;
  display: flex;
  flex-flow: row wrap;
  align-items: center;
  font-size: 0.9rem;
  font-weight: bold;
  padding: 0.5rem 0.25rem 0.5rem 0.25rem;
  cursor: pointer;
  color: ${themePalette.text1};
`;

export default UserMenuButtonWebNavigation;
