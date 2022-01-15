import { css } from '@emotion/react';
import { usePublicNavigationValue } from '../../../atoms/publicNavigationState';
import PublicNavigationButton from './PublicNavigationButton';
import PublicNavigationLinks from './PublicNavigationLinks';

interface PublicNavigationProps {}

function PublicNavigation(props: PublicNavigationProps) {
  const { visible } = usePublicNavigationValue();

  return (
    <div css={block}>
      <PublicNavigationButton />
      <PublicNavigationLinks visible={visible} />
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

export default PublicNavigation;
