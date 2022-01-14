import { css } from '@emotion/react';
import { memo } from 'react';
import { NavLink } from 'react-router-dom';
import {
  useHomeNavigationActions,
  useHomeNavigationValue,
} from '../../../atoms/homeNavigationState';
import { MenuIcon } from '../../../image/icons';
import useTransitionTimeoutEffect from '../../../lib/hooks/useTransitionTimeoutEffect';
import { palette, themeColor, transitions, zIndexes } from '../../../styles';

const NavigationLink: React.FC<{ to: string; text: string }> = ({
  to,
  text,
}) => {
  return (
    <li css={listBlock}>
      <NavLink css={link} to={to}>
        {text}
      </NavLink>
    </li>
  );
};

const link = css`
  display: flex;
  flex-flow: row wrap;
  align-items: center;
  font-size: 1rem;
  padding: 0.5rem 0.25rem 0.5rem 0.25rem;
  &.active {
    font-weight: bold;
  }
`;

const Navigation: React.FC<{ visible: boolean }> = memo(({ visible }) => {
  const closed = useTransitionTimeoutEffect({ visible, delay: 125 });

  if (!visible && closed) return null;

  return (
    <nav css={navigationBlock(visible)}>
      <ul css={listBlock}>
        <NavigationLink to="/" text="홈" />
        <NavigationLink to="/preview" text="미리보기" />
        <NavigationLink to="/guides" text="가이드 및 튜토리얼" />
        <NavigationLink to="/notice" text="공지사항" />
        <NavigationLink to="/policy" text="서비스 정책" />
      </ul>
    </nav>
  );
});

const navigationBlock = (visible: boolean) => css`
  position: absolute;
  top: 3.25rem;
  left: 0.5rem;
  width: 15.15rem;
  height: auto;
  padding: 0.25rem 0.5rem;
  display: flex;
  flex-direction: column;
  z-index: ${zIndexes.dropdownItem};
  border-radius: 0.25rem;
  background: ${themeColor.navigation['light']};
  box-shadow: ${themeColor.shadow['light']};
  transform-origin: top;
  ${visible
    ? css`
        animation: ${transitions.scaleUp} 0.125s forwards ease-in-out;
      `
    : css`
        animation: ${transitions.scaleDown} 0.125s forwards ease-in-out;
      `};
`;

const listBlock = css`
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
`;

interface HomeNavigationProps {}

function HomeNavigation(props: HomeNavigationProps) {
  const { visible } = useHomeNavigationValue();
  const { handleNavigation } = useHomeNavigationActions();

  return (
    <div css={block}>
      <button css={button} onClick={handleNavigation}>
        <MenuIcon />
      </button>
      <Navigation visible={visible} />
    </div>
  );
}

const block = css`
  display: flex;
  justify-content: center;
  align-items: center;
  &:nth-of-type(1) {
    padding: 0 0.5rem;
  }
  &:nth-of-type(2) {
    padding-left: 0.5rem;
  }
`;

const button = css`
  background: inherit;
  border: none;
  box-shadow: none;
  border-radius: 0;
  overflow: visible;
  cursor: pointer;
  display: flex;
  align-items: center;
  padding: 0.5rem;
  border-radius: 3px;
  color: ${palette.gray['500']};
  svg {
    width: 1.125rem;
    height: 1.125rem;
    fill: ${themeColor.fill['light']};
  }
  &:hover {
    svg {
      opacity: 0.5;
    }
  }
`;

export default HomeNavigation;
