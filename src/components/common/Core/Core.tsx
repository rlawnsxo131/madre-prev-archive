import GlobalStyle from '../GlobalStyle';
import { PopupCommon } from '../Popup';

interface CoreProps {}

function Core(props: CoreProps) {
  return (
    <>
      <PopupCommon />
      {/* <div>toast</div> */}
      <GlobalStyle />
    </>
  );
}

export default Core;
