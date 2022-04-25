import { css } from '@emotion/react';
import { zIndexes } from '../../../styles';
import HeaderWeb from './HeaderWeb';
import HeaderMobile from './HeaderMobile';

interface HeaderProps {}

function Header(props: HeaderProps) {
  return (
    <header css={block}>
      <HeaderWeb />
      <HeaderMobile />
    </header>
  );
}

const block = css`
  position: sticky;
  top: 0;
  left: 0;
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 0.25rem 0;
  z-index: ${zIndexes.layoutHeader};
`;

export default Header;
