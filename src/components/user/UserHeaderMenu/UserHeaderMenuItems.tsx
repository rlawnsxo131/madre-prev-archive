import { css } from '@emotion/react';
import useTransitionTimeoutEffect from '../../../hooks/useTransitionTimeoutEffect';
import { themePalette, transitions, zIndexes } from '../../../styles';
import MadreLink from '../../common/MadreLink';

interface UserHeaderMenuItemsProps {
  signOut: () => Promise<void>;
  visible: boolean;
  display_name: string;
}

function UserHeaderMenuItems({
  signOut,
  visible,
  display_name,
}: UserHeaderMenuItemsProps) {
  const closed = useTransitionTimeoutEffect({ visible });

  if (!visible && closed) return null;

  return (
    <div css={block(visible)}>
      <ul css={ul}>
        <li>
          <MadreLink
            to={`/@${display_name}`}
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
  top: 3.25rem;
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

export default UserHeaderMenuItems;
