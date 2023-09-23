import { css } from '@emotion/react';
import { Link } from 'react-router-dom';

interface HeaderLogoProps {}

function HeaderLogo(props: HeaderLogoProps) {
  return (
    <div css={block}>
      <Link css={link} to="/">
        <h1>Madre</h1>
      </Link>
    </div>
  );
}

const block = css`
  display: flex;
  justify-content: center;
  align-items: center;
`;

const link = css`
  display: inline-flex;
  flex-flow: row wrap;
  font-size: 1.25rem;
  h1 {
    margin: 0;
    padding: 0;
    font-weight: bold;
    font-size: 1.5rem;
  }
`;

export default HeaderLogo;
