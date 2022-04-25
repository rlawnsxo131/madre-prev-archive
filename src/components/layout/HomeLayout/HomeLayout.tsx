import { css } from '@emotion/react';
import { Outlet } from 'react-router-dom';
import Header from '../../layout/Header';
import layoutStyles from '../layout.styles';

interface HomeTemplateProps {}

function HomeTemplate(props: HomeTemplateProps) {
  return (
    <div css={block}>
      <Header />
      <main css={main}>
        <div css={content}>
          <Outlet />
        </div>
      </main>
    </div>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const main = css`
  display: flex;
  justify-content: center;
  position: relative;
  ${layoutStyles.responsive}
`;

const content = css`
  display: flex;
  flex-direction: column;
  width: 100%;
`;

export default HomeTemplate;
