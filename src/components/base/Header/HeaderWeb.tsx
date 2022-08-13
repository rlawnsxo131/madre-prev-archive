import { css } from '@emotion/react';
import { media, mediaQuery } from '../../../styles';

interface HeaderWebProps {
  children: React.ReactNode;
  logo: React.ReactNode;
  navigation: React.ReactNode;
}

function HeaderWeb({ children, logo, navigation }: HeaderWebProps) {
  return (
    <div css={block}>
      <div css={itemBlock}>
        {logo}
        {navigation}
      </div>
      <div css={iconItemsBlock}>{children}</div>
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

const itemBlock = css`
  display: flex;
  justify-content: center;
  gap: 1.5rem;
`;

const iconItemsBlock = css`
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1.5rem;
  padding: 0 1rem;
`;

export default HeaderWeb;
