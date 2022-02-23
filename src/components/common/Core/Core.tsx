import GlobalStyle from '../GlobalStyle';
import PopupCommon from '../PopupCommon';
import PopupAuth from '../PopupAuth';

interface CoreProps {}

function Core(props: CoreProps) {
  return (
    <>
      <GlobalStyle />
      <PopupCommon />
      <PopupAuth />
    </>
  );
}

export default Core;
