import { memo } from 'react';
import { css } from '@emotion/react';
import useTransitionTimeoutEffect from '../../../hooks/useTransitionTimeoutEffect';
import {
  basicStyles,
  themePalette,
  transitions,
  zIndexes,
} from '../../../styles';
import MadreLink from '../../common/MadreLink';

interface HeaderUserMenuItemsProps {
  signOut: () => Promise<void>;
  visible: boolean;
  username: string;
}

function HeaderUserMenuItems({
  signOut,
  visible,
  username,
}: HeaderUserMenuItemsProps) {
  const closed = useTransitionTimeoutEffect({ visible });

  if (!visible && closed) return null;

  return (
    <div css={block(visible)}>
      <ul css={ul}>
        <li>
          <MadreLink
            to={`/@${username}`}
            displayName="마이 페이지"
            parentDirection="column"
          />
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
  top: 3rem;
  right: 0;
  width: 10rem;
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

const button = css`
  ${basicStyles.button};
  display: flex;
  flex-flow: row wrap;
  align-items: center;
  font-size: 0.9rem;
  font-weight: bold;
  padding: 0.5rem 0.25rem 0.5rem 0.25rem;
  color: ${themePalette.text1};
`;

export default memo(HeaderUserMenuItems);
