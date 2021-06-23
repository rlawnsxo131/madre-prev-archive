import { css } from '@emotion/react';
import { Link } from 'react-router-dom';
import HomeTemplate from './HomeTemplate';

interface HomeProps {}

function Home(props: HomeProps) {
  return (
    <HomeTemplate>
      <Link to="/user/name">click</Link>
      <div css={block}>home main content</div>
    </HomeTemplate>
  );
}

const block = css``;

export default Home;
