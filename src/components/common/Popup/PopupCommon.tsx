import { css } from '@emotion/react';
import {
  usePopupCommonAction,
  usePopupCommonState,
} from '../../../atoms/popupCommonState';
import media, { mediaQuery } from '../../../styles/media';
import Button from '../Button';
import PopupBase from '../PopupBase';

interface PopupCommonProps {}

function PopupCommon(props: PopupCommonProps) {
  const state = usePopupCommonState();
  const { closePopup } = usePopupCommonAction();
  return (
    <PopupBase visible={state.visible}>
      <div css={block}>
        <h3>{state.title}</h3>
        <p>{state.message}</p>
        <div css={buttonBlock}>
          <Button color="red" onClick={closePopup}>
            확인
          </Button>
        </div>
      </div>
    </PopupBase>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
  padding: 1.725rem 1.5rem;
  h3 {
    margin: 0;
  }
  p {
    padding: 1rem 0;
    line-height: 1.75;
    word-break: break-word;
    overflow-wrap: break-word;
    display: -webkit-box;
  }
  ${media.xxxsmall} {
    width: calc(100vw);
  }
  ${mediaQuery(400)} {
    width: 400px;
  }
`;

const buttonBlock = css`
  display: flex;
  justify-content: flex-end;
  align-items: center;
`;

export default PopupCommon;
