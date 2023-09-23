import { css } from '@emotion/react';
import { media, mediaQuery } from '../../../styles';

interface HeaderWebProps {
  leftSideItems: React.ReactNode;
  rightSideItems: React.ReactNode;
}

function HeaderWeb({ leftSideItems, rightSideItems }: HeaderWebProps) {
  return (
    <div css={block}>
      <div css={leftSideItemsBlock}>{leftSideItems}</div>
      <div css={rightSideItemsBlock}>{rightSideItems}</div>
    </div>
  );
}

const block = css`
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 3.25rem;
  ${media.xxxsmall} {
    display: none;
  }
  ${media.small} {
    width: 93%;
    display: flex;
  }
  ${media.medium} {
    width: 96%;
  }
  ${mediaQuery(1285)} {
    max-width: 1250px;
  }
`;

const leftSideItemsBlock = css`
  display: flex;
  justify-content: center;
  gap: 1.5rem;
`;

const rightSideItemsBlock = css`
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1.5rem;
  padding: 0 1rem;
`;

export default HeaderWeb;
