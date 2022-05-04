import { css } from '@emotion/react';
import useHeaderActions from '../../../hooks/layout/useLayoutHeaderActions';
import { MenuIcon } from '../../../image/icons';
import { palette } from '../../../styles';
import { themePalette } from '../../../styles';

interface HeaderMobileNavigationButtonProps {}

function HeaderMobileNavigationButton(
  props: HeaderMobileNavigationButtonProps,
) {
  const { handleMobileNavigation } = useHeaderActions();

  return (
    <button css={block} onClick={handleMobileNavigation}>
      <MenuIcon />
    </button>
  );
}

const block = css`
  background: inherit;
  border: none;
  box-shadow: none;
  border-radius: 0;
  overflow: visible;
  cursor: pointer;
  display: flex;
  align-items: center;
  padding: 0.5rem 0;
  border-radius: 3px;
  color: ${palette.gray['500']};
  svg {
    width: 1.125rem;
    height: 1.125rem;
    fill: ${themePalette.fill1};
  }
  &:hover {
    svg {
      opacity: 0.5;
    }
  }
`;

export default HeaderMobileNavigationButton;
