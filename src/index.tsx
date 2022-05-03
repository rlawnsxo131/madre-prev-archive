import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';
import { store } from './store';
import App from './App';
import { Storage } from './lib/storage';
import { MADRE_COLOR_THEME } from './constants';
import theme from './store/theme';

const loadTheme = () => {
  const currentTheme = Storage.getItem(MADRE_COLOR_THEME);
  if (!currentTheme) return;
  store.dispatch(theme.actions.setTheme(currentTheme));
  document.body.dataset.theme = currentTheme;
};

loadTheme();

const container = document.getElementById('root')!;
const root = createRoot(container);

root.render(
  <StrictMode>
    <Provider store={store}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </Provider>
  </StrictMode>,
);
