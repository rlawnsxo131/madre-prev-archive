import { renderHook } from '@testing-library/react-hooks';
import useLoadingActions from '../../common/useLoadingActions';

describe('useLoadingActions', () => {
  it('show is called useDispatch', () => {
    const { result } = renderHook(() => useLoadingActions());
  });
});
