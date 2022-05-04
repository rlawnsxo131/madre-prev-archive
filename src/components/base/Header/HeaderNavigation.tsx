import { css } from '@emotion/react';
import { appDisplayRoutes } from '../../../constants';
import MadreLink from '../../common/MadreLink';

interface HeaderNavigationProps {}

function HeaderNavigation(props: HeaderNavigationProps) {
  return (
    <nav css={block}>
      <ul css={ul}>
        {appDisplayRoutes.map((v) => (
          <li key={`app_root_route_${v.path}`}>
            <MadreLink to={v.path} displayName={v.displayName} />
          </li>
        ))}
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

export default HeaderNavigation;
