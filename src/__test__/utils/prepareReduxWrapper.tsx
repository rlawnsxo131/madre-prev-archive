import { configureStore } from '@reduxjs/toolkit';
import { Provider } from 'react-redux';
import { RootReducer } from '../../store';

export default function prepareReduxWrapper(reducer: RootReducer) {
  const store = configureStore({
    reducer,
  });

  const wrapper = ({ children }: { children: React.ReactNode }) => {
    return <Provider store={store}>{children}</Provider>;
  };

  return { wrapper, store } as const;
}
