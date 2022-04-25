import { css } from '@emotion/react';
import { NavLink } from 'react-router-dom';
import { standardColor } from '../../../styles';

interface HeaderNavigationProps {}

function HeaderNavigation(props: HeaderNavigationProps) {
  return (
    <nav css={block}>
      <ul css={ul}>
        <li>
          <NavLink css={link} to="/notice">
            공지사항
          </NavLink>
        </li>
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
      </ul>
    </nav>
  );
}

const block = css`
  display: flex;
  justify-content: space-between;
  padding: 0 1rem;
`;

const ul = css`
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  justify-content: center;
  align-items: center;
`;

const link = css`
  display: flex;
  align-items: center;
  font-size: 0.875rem;
  font-weight: bold;
  padding: 0 1rem;
  &.active {
    color: ${standardColor.navigation.active};
  }
`;

export default HeaderNavigation;
