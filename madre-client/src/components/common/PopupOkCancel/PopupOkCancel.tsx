import { css } from '@emotion/react';
import { media, mediaQuery } from '../../../styles';
import Button from '../Button';
import PopupBase from '../PopupBase';

interface PopupOkCancelProps {
  visible: boolean;
  title?: string;
  message: string;
  onConfirm: () => void | Promise<void>;
  onCancel?: () => void | Promise<void>;
}

function PopupOkCancel({
  visible,
  title,
  message,
  onConfirm,
  onCancel,
}: PopupOkCancelProps) {
  return (
    <PopupBase visible={visible}>
      <div css={block}>
        {title && <h3>{title}</h3>}
        <p>{message}</p>
        <div css={buttonBlock}>
          {onCancel && (
            <Button color="blue" outline onClick={onCancel}>
              취소
            </Button>
          )}
          <Button color="blue" onClick={onConfirm}>
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
    font-size: 1.5rem;
  }
  p {
    margin: 1rem 0;
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
  gap: 0.755rem;
  margin-top: 2rem;
`;

export default PopupOkCancel;
