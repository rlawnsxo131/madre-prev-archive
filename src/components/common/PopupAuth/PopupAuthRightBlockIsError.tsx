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
        <h1 className="is-error">Sorry, An Error Occurred</h1>
      </div>
      <div css={PopupAuthStyles.rightBlockBody}>
        <CancelImage />
      </div>
      <div css={buttonBlock}>
        <Button size="responsive" color="blue" outline onClick={resetError}>
          Try again
        </Button>
        <Button size="responsive" color="blue" onClick={reset}>
          Close
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
