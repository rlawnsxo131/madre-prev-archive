import { css } from '@emotion/react';
import { media, themePalette } from '../../../styles';

interface FooterMobileProps {
  children: React.ReactNode;
}

function FooterMobile({ children }: FooterMobileProps) {
  return (
    <div css={block}>
      <div css={menuBlock}>{children}</div>
      <div css={fakeBlock} />
    </div>
  );
}

const block = css`
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 1rem 1rem 0 0;
  border-top: 1px solid ${themePalette.border_element1};
  ${media.small} {
    display: none;
  }
`;

const menuBlock = css`
  width: 93%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 3rem;
`;

const fakeBlock = css`
  width: 100%;
  height: 1.5rem;
`;

export default FooterMobile;
