import { css } from '@emotion/react';
import { UserIcon } from '../../../image/icons';
import { basicStyles, themePalette } from '../../../styles';
import useFooterMobileUserMenu from './hooks/useFooterMobileUserMenu';

interface FooterMobileUserMenuProps {}

function FooterMobileUserMenu(props: FooterMobileUserMenuProps) {
  const { isActive, onClick } = useFooterMobileUserMenu();

  return (
    <button css={[basicStyles.button, button(isActive)]} onClick={onClick}>
      <UserIcon />
    </button>
  );
}

const button = (isActive?: boolean) => css`
  display: flex;
  justify-content: center;
  align-items: center;
  ${isActive &&
  css`
    svg {
      fill: ${themePalette.anchor_active1};
    }
  `}
`;

export default FooterMobileUserMenu;
