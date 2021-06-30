import { Route, Switch } from 'react-router-dom';
import Core from './components/common/Core';
import GlobalStyle from './components/common/GlobalStyle';
import ErrorBoundary from './components/error';
import HomeTemplate from './components/home/HomeTemplate';
import HomePage from './pages/HomePage';
import NotFoundPage from './pages/NotFoundPage';
import ThreeTestPage from './pages/ThreeTestPage';
import UserPage from './pages/UserPage';

interface AppProps {}

function App(props: AppProps) {
  return (
    <ErrorBoundary>
      <Switch>
        <Route exact path={['/', '/guides', '/notice', '/policy', '/user']}>
          <HomeTemplate>
            <Route exact path="/">
              <HomePage />
            </Route>
            <Route path="/guides">
              <div>guides</div>
            </Route>
            <Route path="/notice">
              <div>notice</div>
            </Route>
            <Route path="/policy">
              <div>policy</div>
            </Route>
            <Route path="/user">
              <UserPage />
            </Route>
          </HomeTemplate>
        </Route>
        <Route path="/three">
          <ThreeTestPage />
        </Route>
        <Route>
          <NotFoundPage />
        </Route>
      </Switch>
      <Core />
      <GlobalStyle />
    </ErrorBoundary>
  );
}

export default App;
