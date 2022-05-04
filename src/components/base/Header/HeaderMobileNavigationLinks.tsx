import { css } from '@emotion/react';
import { transitions, zIndexes } from '../../../styles';
import useTransitionTimeoutEffect from '../../../hooks/useTransitionTimeoutEffect';
import useHeaderState from '../../../hooks/layout/useLayoutHeaderState';
import { themePalette } from '../../../styles';
import { appDisplayRoutes } from '../../../constants';
import MadreLink from '../../common/MadreLink';

interface HeaderMobileNavigationLinksProps {}

function HeaderMobileNavigationLinks(props: HeaderMobileNavigationLinksProps) {
  const {
    navigation: { visible },
  } = useHeaderState();
  const closed = useTransitionTimeoutEffect({ visible });

  if (!visible && closed) return null;

  return (
    <nav css={block(visible)}>
      <ul css={ul}>
        {appDisplayRoutes.map((v) => (
          <li key={`root_route_mobile_${v.path}`}>
            <MadreLink
              to={v.path}
              displayName={v.displayName}
              parentDirection="column"
            />
          </li>
        ))}
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

export default HeaderMobileNavigationLinks;
