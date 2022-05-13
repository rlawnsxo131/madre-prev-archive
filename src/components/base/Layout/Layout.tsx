import { css } from '@emotion/react';
import Footer from '../Footer';
import Header from '../Header';
import baseStyles from '../base.styles';
import { Outlet } from 'react-router-dom';

interface LayoutProps {}

function Layout(props: LayoutProps) {
  return (
    <div css={block}>
      <Header />
      <main css={[baseStyles.responsive, main]}>
        <div css={content}>
          <Outlet />
        </div>
      </main>
      {/* currently, only mobile is supposed to draw footer. */}
      <Footer />
    </div>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100%;
  overflow-y: scroll;
`;

const main = css`
  flex: 1;
  display: flex;
  justify-content: center;
  position: relative;
`;

const content = css`
  display: flex;
  flex-direction: column;
  width: 100%;
`;

export default Layout;
