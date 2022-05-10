import { css } from '@emotion/react';
import { zIndexes } from '../../../styles';
import FooterMobile from './FooterMobile';

interface FooterProps {}

function Footer(props: FooterProps) {
  return (
    <footer css={block}>
      <FooterMobile />
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
  align-items: center;
  z-index: ${zIndexes.layoutFooter};
`;

export default Footer;
