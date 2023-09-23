import { renderHook } from '@testing-library/react-hooks';
import { createElement } from 'react';
import renderWithProviders from '../../../__tests__/utils/renderWithProviders';
import useLoadingActions from '../../common/useLoadingActions';

describe('useLoadingActions', () => {
  it('show is called useDispatch', () => {
    const { store } = renderWithProviders(createElement('div'));
    const { result } = renderHook(() => useLoadingActions());
    console.log('result:', result);
    console.log('store: ', store);
  });
});
