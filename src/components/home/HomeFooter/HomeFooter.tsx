import { css } from '@emotion/react';
import homeStyles from '../home.styles';

interface HomeFooterProps {}

function HomeFooter(props: HomeFooterProps) {
  return <footer css={block}>home footer</footer>;
}

const block = css`
  ${homeStyles.responsive}
`;

export default HomeFooter;
