import { css } from '@emotion/react';
import usePopupAuthActions from '../../../hooks/popupAuth/usePopupAuthActions';
import usePopupAuthState from '../../../hooks/popupAuth/usePopupAuthState';
import { CloseIcon, WarningIcon } from '../../../image/icons';
import { MobileUserImage } from '../../../image/images';
import { InspirationImage, CancelImage } from '../../../image/images';
import { mediaQuery, palette, transitions } from '../../../styles';
import PopupBase from '../PopupBase';
import ButtonGoogleSignin from '../ButtonGoogleSignin';
import Button from '../Button';

interface RightBlockProps {
  close: () => void;
}

function RightBlock({ close }: RightBlockProps) {
  return (
    <div css={rightBlock}>
      <div css={rightBlockHeader}>
        <CloseIcon className="popup-auth-close-icon" onClick={close} />
        <h1>Welcome To Madre</h1>
      </div>
      <div css={rightBlockBody}>
        <InspirationImage />
      </div>
      <ButtonGoogleSignin />
    </div>
  );
}

interface IsErrorRightBlockProps {
  close: () => void;
  resetError: () => void;
}

function IsErrorRightBlock({ close, resetError }: IsErrorRightBlockProps) {
  const reset = () => {
    close();
    setTimeout(() => {
      resetError();
    }, 250);
  };
  return (
    <div css={rightBlock}>
      <div css={rightBlockHeader}>
        <div
          css={css`
            display: flex;
            justify-content: space-between;
          `}
        >
          <WarningIcon fill={palette.pink['600']} />
          <CloseIcon className="popup-auth-close-icon" onClick={reset} />
        </div>
        <h1 className="is-error">Sorry, An Error Occurred!</h1>
      </div>
      <div css={rightBlockBody}>
        <CancelImage />
      </div>
      <div
        css={css`
          display: flex;
          flex-direction: column;
          button {
            margin-top: 1rem;
          }
        `}
      >
        <Button size="responsive" color="pink" outline onClick={resetError}>
          Try again
        </Button>
        <Button size="responsive" color="pink" onClick={reset}>
          Close
        </Button>
      </div>
    </div>
  );
}

interface PopupAuthProps {}

function PopupAuth(props: PopupAuthProps) {
  const { visible, isError } = usePopupAuthState();
  const { close, resetError } = usePopupAuthActions();

  return (
    <PopupBase visible={visible}>
      <div css={block(isError)}>
        <div css={leftBlock}>
          <MobileUserImage />
        </div>
        {isError && <IsErrorRightBlock close={close} resetError={resetError} />}
        {!isError && <RightBlock close={close} />}
      </div>
    </PopupBase>
  );
}

const shakeAnimation = css`
  animation: ${transitions.shake} 0.5s 0.25s ease-in-out;
`;

const block = (isError: boolean) => css`
  display: flex;
  border-radius: 1rem;
  background: white;
  ${mediaQuery(512)} {
    width: 30rem;
  }
  ${isError && shakeAnimation};
`;

const leftBlock = css`
  flex: 1 1 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  background: ${palette.gray['100']};
  border-radius: 0.25rem;
  svg {
    width: 100%;
  }
`;

const rightBlock = css`
  flex: 2 1 0;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 1rem 1rem 1.5rem 1rem;
  border-radius: 0.25rem;
`;

const rightBlockHeader = css`
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 1rem;

  h1 {
    margin: 0;
    &.is-error {
      color: ${palette.pink['600']};
      ${shakeAnimation}
    }
  }

  .popup-auth-close-icon {
    align-self: flex-end;
    width: 1.25rem;
    height: 1.25rem;
    color: ${palette.gray['600']};
    cursor: pointer;
  }
`;

const rightBlockBody = css`
  display: flex;
  flex-direction: column;
  svg {
    width: 100%;
    height: auto;
  }
`;

export default PopupAuth;
