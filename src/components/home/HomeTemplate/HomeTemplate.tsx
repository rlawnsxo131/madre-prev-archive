import { css } from '@emotion/react';
import { Outlet } from 'react-router-dom';
import homeStyles from '../home.styles';
import HomeHeader from '../HomeHeader';

interface HomeTemplateProps {}

function HomeTemplate(props: HomeTemplateProps) {
  return (
    <div css={block}>
      <HomeHeader />
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
  ${homeStyles.responsive}
`;

const content = css`
  display: flex;
  flex-direction: column;
  width: 100%;
`;

export default HomeTemplate;
