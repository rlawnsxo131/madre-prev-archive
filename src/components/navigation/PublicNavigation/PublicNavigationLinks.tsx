import { css } from '@emotion/react';
import { NavLink } from 'react-router-dom';
import { navigationActiveColor } from '../navigation.styles';

interface PublicNavigationLinksProps {}

function PublicNavigationLinks(props: PublicNavigationLinksProps) {
  return (
    <ul css={block}>
      <li>
        <NavLink css={link} to="/preview">
          미리보기
        </NavLink>
      </li>
      <li>
        <NavLink css={link} to="/guides">
          가이드 및 튜토리얼
        </NavLink>
      </li>
    </ul>
  );
}

const block = css`
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
    color: ${navigationActiveColor};
  }
`;

export default PublicNavigationLinks;
