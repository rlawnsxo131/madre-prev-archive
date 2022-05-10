import { css } from '@emotion/react';
import { HomeIcon, MenuIcon } from '../../../image/icons';
import MadreImageLink from '../../common/MadreImageLink';
import FooterMobileNotification from './FooterMobileNotification';
import FooterMobileUserMenu from './FooterMobileUserMenu';

interface FooterMobileItemsProps {}

// TODO: need path change
function FooterMobileItems(props: FooterMobileItemsProps) {
  return (
    <div css={block}>
      <MadreImageLink to="/">
        <HomeIcon />
      </MadreImageLink>
      <MadreImageLink to="/@name">
        <FooterMobileUserMenu />
      </MadreImageLink>
      <MadreImageLink to="/notifications">
        <FooterMobileNotification />
      </MadreImageLink>
      <MadreImageLink to="/">
        <MenuIcon />
      </MadreImageLink>
    </div>
  );
}

const block = css`
  width: 93%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 3rem;
`;

export default FooterMobileItems;
