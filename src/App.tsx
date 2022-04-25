import { Route, Routes } from 'react-router-dom';
import Core from './components/common/Core';
import ErrorBoundary from './components/error/ErrorBoundary';
import HomeLayout from './components/layout/HomeLayout';
import { HomePage, NotFoundPage, TestPage2 } from './pages';

interface AppProps {}

function App(props: AppProps) {
  return (
    <ErrorBoundary>
      <Routes>
        <Route path="/" element={<HomeLayout />}>
          <Route index element={<HomePage />} />
          <Route path="guide" element={<div>guide</div>} />
          <Route path="notice" element={<div>notice</div>} />
          <Route path="policy" element={<div>policy</div>} />
          <Route path="preview" element={<div>preview</div>} />
        </Route>
        <Route path="/test" element={<TestPage2 />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
      <Core />
    </ErrorBoundary>
  );
}

export default App;
