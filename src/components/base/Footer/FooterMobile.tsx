import { css } from '@emotion/react';
import { media } from '../../../styles';

interface FooterMobileProps {}

function FooterMobile(props: FooterMobileProps) {
  return (
    <div css={block}>
      <div>a</div>
      <div>a</div>
      <div>a</div>
      <div>a</div>
    </div>
  );
}

const block = css`
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 3rem;
  padding: 0.25rem 0;

  ${media.xxxsmall} {
    width: 93%;
  }
  ${media.small} {
    display: none;
  }
`;

export default FooterMobile;
