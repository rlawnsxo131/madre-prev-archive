import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import { RecoilRoot } from 'recoil';
import App from './App';
import recoilInitializer from './atoms/recoilInitializer';

ReactDOM.render(
  <React.StrictMode>
    <RecoilRoot initializeState={recoilInitializer}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </RecoilRoot>
  </React.StrictMode>,
  document.getElementById('root'),
);
