import GlobalStyle from '../GlobalStyle';
import PopupCommon from '../PopupCommon';
import PopupAuth from '../PopupAuth';
import ScreenSignUp from '../ScreenSignUp';
import Loading from '../Loading';
import Toast from '../Toast';
import useThemeLoadEffect from '../../../hooks/theme/useThemeLoadEffect';
import useUserLoadEffect from '../../../hooks/user/useUserLoadEffect';

interface CoreProps {}

function Core(props: CoreProps) {
  useThemeLoadEffect();
  useUserLoadEffect();

  return (
    <>
      <GlobalStyle />
      <PopupCommon />
      <PopupAuth />
      <ScreenSignUp />
      <Loading />
      <Toast />
    </>
  );
}

export default Core;
