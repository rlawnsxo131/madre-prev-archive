import { css } from '@emotion/react';
import HeaderMobileNavigationButton from './HeaderMobileNavigationButton';
import HeaderMobileNavigationLinks from './HeaderMobileNavigationLinks';

interface HeaderMobileNavigationProps {}

function HeaderMobileNavigation(props: HeaderMobileNavigationProps) {
  return (
    <div css={block}>
      <HeaderMobileNavigationButton />
      <HeaderMobileNavigationLinks />
    </div>
  );
}

const block = css`
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  &:nth-of-type(1) {
    padding: 0 0.5rem;
  }
  &:nth-of-type(2) {
    padding-left: 0.5rem;
  }
`;

export default HeaderMobileNavigation;
