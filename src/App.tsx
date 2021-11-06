import { Route, Routes } from 'react-router-dom';
import Core from './components/common/Core';
import ErrorBoundary from './components/error';
import HomeTemplate from './components/home/HomeTemplate';
import HomePage from './pages/HomePage';
import NotFoundPage from './pages/NotFoundPage';
import TestPage from './pages/TestPage';

interface AppProps {}

function App(props: AppProps) {
  return (
    <ErrorBoundary>
      <Routes>
        <Route path="/" element={<HomeTemplate />}>
          <Route index element={<HomePage />} />
          <Route path="guides" element={<div>guides</div>} />
          <Route path="notice" element={<div>notice</div>} />
          <Route path="policy" element={<div>policy</div>} />
          <Route path="user" element={<div>user</div>} />
        </Route>
        <Route path="/test" element={<TestPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
      <Core />
    </ErrorBoundary>
  );
}

export default App;
