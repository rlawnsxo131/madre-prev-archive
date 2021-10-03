import { css } from '@emotion/react';
import { memo } from 'react';
import { NavLink } from 'react-router-dom';
import { ColorTheme } from '../../../atoms/colorThemeState';
import { useTransitionTimeoutEffect } from '../../../lib/hooks';
import { themeColor } from '../../../styles/palette';
import transitions from '../../../styles/transitions';
import zIndexes from '../../../styles/zIndexes';

interface HomeNavigationProps {
  theme: ColorTheme;
  visible: boolean;
}

function HomeNavigation({ theme, visible }: HomeNavigationProps) {
  const closed = useTransitionTimeoutEffect({ visible, delay: 125 });

  if (!visible && closed) return null;

  return (
    <nav css={block(theme, visible)}>
      <ul>
        <li>
          <NavLink css={link} exact to="/">
            홈
          </NavLink>
        </li>
        <li>
          <NavLink css={link} to="/guides">
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

const block = (theme: ColorTheme, visible: boolean) => css`
  position: absolute;
  top: 3.25rem;
  left: 0.5rem;
  width: 18rem;
  height: auto;
  padding: 0.25rem 0.5rem;
  display: flex;
  flex-direction: column;
  z-index: ${zIndexes.dropdownItem};
  border-radius: 0.25rem;
  background: ${themeColor.background[theme]};
  box-shadow: ${themeColor.shadow[theme]};
  transform-origin: top;
  ul,
  li {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }
  ${visible
    ? css`
        animation: ${transitions.scaleUp} 0.125s forwards ease-in-out;
      `
    : css`
        animation: ${transitions.scaleDown} 0.125s forwards ease-in-out;
      `};
`;

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

export default memo(HomeNavigation);
