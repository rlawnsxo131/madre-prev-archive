import { Route, Routes } from 'react-router-dom';
import Core from './components/common/Core';
import ErrorBoundary from './components/error/ErrorBoundary';
import Layout from './components/base/Layout';
import { HomePage, MadreStoryPage, NotFoundPage } from './pages';
import UserPage from './pages/UserPage';

interface AppProps {}

function App(props: AppProps) {
  return (
    <ErrorBoundary>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<HomePage />} />
          <Route path="madre-story" element={<MadreStoryPage />} />
          <Route path="notice" element={<div>notice</div>} />
          <Route path="guide" element={<div>guide</div>} />
          <Route path="policy" element={<div>policy</div>} />
          <Route path="notifications" element={<div>notifications</div>} />
        </Route>
        <Route path="/@:username" element={<Layout />}>
          <Route index element={<UserPage />} />
          <Route path="profile" element={<div>user profile</div>} />
        </Route>
        <Route path="/m" element={<Layout />}>
          <Route path="all-menu" element={<div>mobile all menu</div>} />
        </Route>
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
      <Core />
    </ErrorBoundary>
  );
}

export default App;
