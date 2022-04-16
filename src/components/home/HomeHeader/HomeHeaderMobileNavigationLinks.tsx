import { css } from '@emotion/react';
import { NavLink } from 'react-router-dom';
import { standardColor, transitions, zIndexes } from '../../../styles';
import useTransitionTimeoutEffect from '../../../hooks/useTransitionTimeoutEffect';
import useHomeHeaderState from '../../../hooks/home/useHomeHeaderState';
import { themePalette } from '../../../styles';

interface HomeHeaderMobileNavigationLinksProps {}

function HomeHeaderMobileNavigationLinks(
  props: HomeHeaderMobileNavigationLinksProps,
) {
  const {
    navigation: { visible },
  } = useHomeHeaderState();
  const closed = useTransitionTimeoutEffect({ visible });

  if (!visible && closed) return null;

  return (
    <nav css={block(visible)}>
      <ul css={ul}>
        <li>
          <NavLink css={link} to="/preview">
            미리보기
          </NavLink>
        </li>
        <li>
          <NavLink css={link} to="/guide">
            가이드 및 튜토리얼
          </NavLink>
        </li>
        <li>
          <NavLink css={link} to="/notice">
            공지사항
          </NavLink>
        </li>
        <li>
          <NavLink css={link} to="/policy">
            서비스 정책
          </NavLink>
        </li>
      </ul>
    </nav>
  );
}

const block = (visible: boolean) => css`
  position: absolute;
  top: 3.25rem;
  right: 0;
  width: 12rem;
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

export default HomeHeaderMobileNavigationLinks;
