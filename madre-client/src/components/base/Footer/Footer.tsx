import { css } from '@emotion/react';
import { HomeIcon, MenuIcon } from '../../../image/icons';
import { zIndexes } from '../../../styles';
import LinkImage from '../../common/LinkImage';
import FooterMobile from './FooterMobile';
import FooterMobileNotification from './FooterMobileNotification';
import FooterMobileUserMenu from './FooterMobileUserMenu';
import useFooterMobileUserMenu from './hooks/useFooterMobileUserMenu';

interface FooterProps {}

function Footer(props: FooterProps) {
  const { isActive, onClick } = useFooterMobileUserMenu();

  return (
    <footer css={block}>
      <FooterMobile>
        <LinkImage to="/">
          <HomeIcon />
        </LinkImage>
        <FooterMobileUserMenu isActive={isActive} onClick={onClick} />
        <LinkImage to="/notifications">
          <FooterMobileNotification />
        </LinkImage>
        <LinkImage to="/m/all-menu">
          <MenuIcon />
        </LinkImage>
      </FooterMobile>
    </footer>
  );
}

const block = css`
  position: sticky;
  bottom: 0;
  left: 0;
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  z-index: ${zIndexes.layoutFooter};
`;

export default Footer;
