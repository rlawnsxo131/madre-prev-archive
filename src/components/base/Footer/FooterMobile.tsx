import { css } from '@emotion/react';
import { media, themePalette } from '../../../styles';
import FooterMobileItems from './FooterMobileItems';

interface FooterMobileProps {}

function FooterMobile(props: FooterMobileProps) {
  return (
    <div css={block}>
      <FooterMobileItems />
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

const fakeBlock = css`
  width: 100%;
  height: 1.5rem;
`;

export default FooterMobile;
