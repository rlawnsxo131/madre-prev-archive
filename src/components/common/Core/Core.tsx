import GlobalStyle from '../GlobalStyle';
import PopupCommon from '../PopupCommon';
import PopupLogin from '../PopupLogin';

interface CoreProps {}

function Core(props: CoreProps) {
  return (
    <>
      <GlobalStyle />
      <PopupCommon />
      <PopupLogin />
      {/* <div>toast</div> */}
    </>
  );
}

export default Core;
