import { css } from '@emotion/react';
import { media } from '../../../styles';

interface HeaderMobileProps {
  children: React.ReactNode;
  logo: React.ReactNode;
}

function HeaderMobile({ logo, children }: HeaderMobileProps) {
  return (
    <div css={block}>
      {logo}
      {children}
    </div>
  );
}

const block = css`
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 3rem;

  ${media.xxxsmall} {
    width: 93%;
  }
  ${media.small} {
    display: none;
  }
`;

export default HeaderMobile;
