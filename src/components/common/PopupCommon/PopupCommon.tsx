import { css } from '@emotion/react';
import { media, mediaQuery } from '../../../styles';
import Button from '../Button';
import PopupBase from '../PopupBase';
import usePopupCommonState from '../../../hooks/common/usePopupCommonState';
import usePopupCommonActions from '../../../hooks/common/usePopupCommonActions';

interface PopupCommonProps {}

function PopupCommon(props: PopupCommonProps) {
  const { close } = usePopupCommonActions();
  const { visible, title, message } = usePopupCommonState();

  return (
    <PopupBase visible={visible}>
      <div css={block}>
        {title && <h3>{title}</h3>}
        <p>{message}</p>
        <div css={buttonBlock}>
          <Button color="blue" onClick={close}>
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
  padding: 2rem 1.5rem;

  h3 {
    margin: 0;
  }
  p {
    margin: 0;
    padding: 1rem 0;
    line-height: 1.75;
    word-break: break-word;
    overflow-wrap: break-word;
    display: -webkit-box;
  }

  ${media.xxxsmall} {
    width: 100vw;
  }
  ${mediaQuery(400)} {
    width: 25rem;
  }
`;

const buttonBlock = css`
  display: flex;
  justify-content: flex-end;
  align-items: center;
`;

export default PopupCommon;
