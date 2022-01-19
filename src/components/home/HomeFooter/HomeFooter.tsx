import { css } from '@emotion/react';
import homeStyles from '../home.styles';

interface HomeFooterProps {}

function HomeFooter(props: HomeFooterProps) {
  return <div css={block}>home footer</div>;
}

const block = css`
  width: 100%;
  ${homeStyles.responsive}
`;

export default HomeFooter;
