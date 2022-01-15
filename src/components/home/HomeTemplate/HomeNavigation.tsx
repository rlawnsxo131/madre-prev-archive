import { css } from '@emotion/react';
import { useHomeNavigationValue } from '../../../atoms/homeNavigationState';
import HomeNavigationButton from './HomeNavigationButton';
import HomeNavigationLinks from './HomeNavigationLinks';

interface HomeNavigationProps {}

function HomeNavigation(props: HomeNavigationProps) {
  const { visible } = useHomeNavigationValue();

  return (
    <div css={block}>
      <HomeNavigationButton />
      <HomeNavigationLinks visible={visible} />
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

export default HomeNavigation;
