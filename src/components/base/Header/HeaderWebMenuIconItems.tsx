import { css } from '@emotion/react';
import ButtonThemeChange from '../../common/ButtonThemeChange';
import HeaderUserMenu from '../HeaderUserMenu';
import HeaderUserNotification from '../HeaderUserNotification';

interface HeaderWebMenuIconItemsProps {}

function HeaderWebMenuIconItems(props: HeaderWebMenuIconItemsProps) {
  return (
    <div css={block}>
      <HeaderUserNotification />
      <HeaderUserMenu />
      <ButtonThemeChange />
    </div>
  );
}

const block = css`
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1.5rem;
  padding: 0 1rem;
`;

export default HeaderWebMenuIconItems;
