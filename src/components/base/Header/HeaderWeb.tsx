import { css } from '@emotion/react';
import { media, mediaQuery } from '../../../styles';
import HeaderLogo from './HeaderLogo';
import HeaderWebNavigation from './HeaderWebNavigation';
import HeaderWebMenuIconItems from './HeaderWebMenuIconItems';

interface HeaderWebProps {}

function HeaderWeb(props: HeaderWebProps) {
  return (
    <div css={block}>
      <HeaderLogo />
      <div css={itemBlock}>
        <HeaderWebNavigation />
        <HeaderWebMenuIconItems />
      </div>
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
`;

export default HeaderWeb;
