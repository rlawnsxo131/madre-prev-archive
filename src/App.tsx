import { css, Global } from '@emotion/react';
import { Route, Switch } from 'react-router-dom';
import Core from './components/common/Core';
import ErrorBoundary from './components/error';
import UserHome from './components/user/UserHome';
import HomePage from './pages/HomePage';

interface AppProps {}

function App(props: AppProps) {
  return (
    <ErrorBoundary>
      <Switch>
        <Route exact path="/">
          <HomePage />
        </Route>
        <Route path="/user/name">
          <UserHome />
        </Route>
        <Route>
          <div>not found</div>
        </Route>
      </Switch>
      <Core />
      <Global styles={globalStyle} />
    </ErrorBoundary>
  );
}

const globalStyle = css`
  html,
  body,
  #root {
    margin: 0;
    padding: 0;
    height: 100%;
  }
  html {
    box-sizing: border-box;
    * {
      box-sizing: inherit;
      font-family: 'Montserrat', sans-serif, -apple-system, BlinkMacSystemFont,
        'Helvetica Neue', 'Apple SD Gothic Neo', 'Malgun Gothic', '맑은 고딕',
        나눔고딕, 'Nanum Gothic', 'Noto Sans KR', 'Noto Sans CJK KR', arial,
        돋움, Dotum, Tahoma, Geneva, sans-serif;
      text-rendering: optimizeLegibility;
      -webkit-font-smoothing: subpixel-antialiased;
      -webkit-font-smoothing: antialiased;
      -moz-osx-font-smoothing: grayscale;
    }
  }
  a {
    text-decoration: none;
    cursor: pointer;
    padding: 0;
    margin: 0;
  }
`;

export default App;
