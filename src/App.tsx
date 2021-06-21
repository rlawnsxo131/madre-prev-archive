import { css, Global } from '@emotion/react';
import { Route, Switch } from 'react-router-dom';
import ErrorBoundary from './components/error';
import Layout from './components/layout';

interface AppProps {}

function App(props: AppProps) {
  return (
    <ErrorBoundary>
      <Layout>
        <Switch>
          <Route exact path="/">
            <div>hello world</div>
          </Route>
        </Switch>
      </Layout>
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
    }
  }
`;

export default App;
