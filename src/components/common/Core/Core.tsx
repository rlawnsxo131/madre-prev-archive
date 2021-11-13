import GlobalStyle from '../GlobalStyle';
import PopupCommon from '../PopupCommon';

interface CoreProps {}

function Core(props: CoreProps) {
  return (
    <>
      <GlobalStyle />
      <PopupCommon />
      {/* <div>toast</div> */}
    </>
  );
}

export default Core;
