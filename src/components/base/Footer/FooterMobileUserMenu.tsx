import { css } from '@emotion/react';
import { UserIcon } from '../../../image/icons';

interface FooterMobileUserMenuProps {}

function FooterMobileUserMenu(props: FooterMobileUserMenuProps) {
  return (
    <div css={block}>
      <UserIcon />
    </div>
  );
}

const block = css`
  display: flex;
  justify-content: center;
  align-items: center;
`;

export default FooterMobileUserMenu;
