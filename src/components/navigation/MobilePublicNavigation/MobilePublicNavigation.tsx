import { css } from '@emotion/react';
import { useMobilePublicNavigationValue } from '../../../atoms/mobilePublicNavigationState';
import { navigationMobileDisplay } from '../navigation.styles';
import MobilePublicNavigationButton from './MobilePublicNavigationButton';
import MobilePublicNavigationLinks from './MobilePublicNavigationLinks';

interface MobilePublicNavigationProps {}

function MobilePublicNavigation(props: MobilePublicNavigationProps) {
  const { visible } = useMobilePublicNavigationValue();

  return (
    <div css={block}>
      <MobilePublicNavigationButton />
      <MobilePublicNavigationLinks visible={visible} />
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
  ${navigationMobileDisplay};
`;

export default MobilePublicNavigation;
