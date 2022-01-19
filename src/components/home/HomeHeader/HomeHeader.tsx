import { css } from '@emotion/react';
import { zIndexes } from '../../../styles';
import HomeHeaderWeb from './HomeHeaderWeb';
import HomeHeaderMobile from './HomeHeaderMobile';

interface HomeHeaderProps {}

function HomeHeader(props: HomeHeaderProps) {
  return (
    <header css={block}>
      <HomeHeaderWeb />
      <HomeHeaderMobile />
    </header>
  );
}

const block = css`
  position: sticky;
  top: 0;
  left: 0;
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 0.25rem 0;
  z-index: ${zIndexes.homeHeader};
`;

export default HomeHeader;
