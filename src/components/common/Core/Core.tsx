import GlobalStyle from '../GlobalStyle';
import PopupCommon from '../PopupCommon';
import PopupAuth from '../PopupAuth';
import ScreenSignup from '../ScreenSignup';
import Loading from '../Loading';

interface CoreProps {}

function Core(props: CoreProps) {
  return (
    <>
      <GlobalStyle />
      <PopupCommon />
      <PopupAuth />
      <ScreenSignup />
      <Loading />
    </>
  );
}

export default Core;
