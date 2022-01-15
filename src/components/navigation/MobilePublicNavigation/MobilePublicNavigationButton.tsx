import { css } from '@emotion/react';
import { useMobilePublicNavigationActions } from '../../../atoms/mobilePublicNavigationState';
import { MenuIcon } from '../../../image/icons';
import { palette, themeColor } from '../../../styles';

interface MobilePublicNavigationButtonProps {}

function MobilePublicNavigationButton(
  props: MobilePublicNavigationButtonProps,
) {
  const { handleNavigation } = useMobilePublicNavigationActions();

  return (
    <button css={block} onClick={handleNavigation}>
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

export default MobilePublicNavigationButton;
