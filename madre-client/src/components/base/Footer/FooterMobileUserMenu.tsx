import { css } from '@emotion/react';
import { memo } from 'react';
import { UserIcon } from '../../../image/icons';
import { basicStyles, themePalette } from '../../../styles';

interface FooterMobileUserMenuProps {
  isActive?: boolean;
  onClick: () => void;
}

function FooterMobileUserMenu({
  isActive,
  onClick,
}: FooterMobileUserMenuProps) {
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

export default memo(FooterMobileUserMenu);
