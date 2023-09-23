import { css } from '@emotion/react';
import { CloseIcon, WarningIcon } from '../../../image/icons';
import { CancelImage } from '../../../image/images';
import { palette } from '../../../styles';
import Button from '../Button';
import PopupAuthStyles from './PopupAuth.styles';

interface PopupAuthRightBlockIsErrorProps {
  close: () => void;
  resetError: () => void;
}

function PopupAuthRightBlockIsError({
  close,
  resetError,
}: PopupAuthRightBlockIsErrorProps) {
  const reset = () => {
    close();
    setTimeout(() => {
      resetError();
    }, 250);
  };

  return (
    <div css={PopupAuthStyles.rightBlock}>
      <div css={PopupAuthStyles.rightBlockHeader}>
        <div css={headerIconBlock}>
          <WarningIcon fill={palette.blue['600']} />
          <CloseIcon className="popup-auth-close-icon" onClick={reset} />
        </div>
        <h1 className="is-error">이런! 오류가 발생했습니다.</h1>
      </div>
      <div css={PopupAuthStyles.rightBlockBody}>
        <CancelImage />
      </div>
      <div css={buttonBlock}>
        <Button size="responsive" color="blue" outline onClick={resetError}>
          다시 시도
        </Button>
        <Button size="responsive" color="blue" onClick={reset}>
          닫기
        </Button>
      </div>
    </div>
  );
}

const headerIconBlock = css`
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const buttonBlock = css`
  display: flex;
  flex-direction: column;
  button {
    margin-top: 1rem;
  }
`;

export default PopupAuthRightBlockIsError;
