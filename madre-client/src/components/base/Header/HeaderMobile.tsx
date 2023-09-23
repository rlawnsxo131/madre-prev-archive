import { css } from '@emotion/react';
import { media } from '../../../styles';

interface HeaderMobileProps {
  leftSideItems: React.ReactNode;
  rightSideItems: React.ReactNode;
}

function HeaderMobile({ leftSideItems, rightSideItems }: HeaderMobileProps) {
  return (
    <div css={block}>
      {leftSideItems}
      {rightSideItems}
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
