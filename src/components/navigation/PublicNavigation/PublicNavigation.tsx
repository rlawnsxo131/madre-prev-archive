import { memo } from 'react';
import { css } from '@emotion/react';
import { useAppNavigationValue } from '../../../atoms/appNavigationState';
import AppNavigationButton from './PublicNavigationButton';
import AppNavigationLinks from './PublicNavigationLinks';

interface AppNavigationProps {}

function PublicNavigation(props: AppNavigationProps) {
  const { visible } = useAppNavigationValue();

  return (
    <div css={block}>
      <AppNavigationButton />
      <AppNavigationLinks visible={visible} />
    </div>
  );
}

const block = css`
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

export default memo(PublicNavigation);
