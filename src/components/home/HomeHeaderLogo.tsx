import { css } from '@emotion/react';
import { Link } from 'react-router-dom';

interface HomeHeaderLogoProps {}

function HomeHeaderLogo(props: HomeHeaderLogoProps) {
  return (
    <div css={block}>
      <Link css={link} to="/">
        <h1>Data Visualizer</h1>
      </Link>
    </div>
  );
}

const block = css``;

const link = css`
  display: inline-flex;
  flex-flow: row wrap;
  font-size: 1.25rem;
  padding: 0.5rem;
  h1 {
    margin: 0;
    padding: 0;
    font-weight: bold;
    font-size: 1.5rem;
  }
`;

export default HomeHeaderLogo;
