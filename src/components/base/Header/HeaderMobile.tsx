import { css } from '@emotion/react';
import { media } from '../../../styles';
import ButtonThemeChange from '../../common/ButtonThemeChange';
import HeaderLogo from './HeaderLogo';

interface HeaderMobileProps {}

function HeaderMobile(props: HeaderMobileProps) {
  return (
    <div css={block}>
      <HeaderLogo />
      <div css={right}>
        <ButtonThemeChange />
      </div>
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

const right = css`
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1.5rem;
`;

export default HeaderMobile;
