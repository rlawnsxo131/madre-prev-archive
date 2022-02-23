import { css } from '@emotion/react';
import usePopupAuthActions from '../../../hooks/core/usePopupAuthActions';
import usePopupAuthState from '../../../hooks/core/usePopupAuthState';
import CloseIcon from '../../../image/icons/CloseIcon';
import { MobileUserImage } from '../../../image/images';
import InspirationImage from '../../../image/images/InspirationImage';
import { mediaQuery, palette } from '../../../styles';
import GoogleAuthButton from '../../auth/GoogleAuthButton';
import PopupBase from '../PopupBase';

interface PopupAuthProps {}

function PopupAuth(props: PopupAuthProps) {
  const { visible } = usePopupAuthState();
  const { onClose } = usePopupAuthActions();

  return (
    <PopupBase visible={visible}>
      <div css={block}>
        <div css={leftBlock}>
          <MobileUserImage />
        </div>
        <div css={rightBlock}>
          <div css={rightBlockHeader}>
            <CloseIcon onClick={onClose} />
            <h1>Welcome To Madre</h1>
          </div>
          <div css={rightBlockBody}>
            <InspirationImage />
          </div>
          <GoogleAuthButton />
        </div>
      </div>
    </PopupBase>
  );
}

const block = css`
  display: flex;
  border-radius: 1rem;
  background: white;
  ${mediaQuery(512)} {
    width: 30rem;
  }
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
  padding: 1rem;
  border-radius: 0.25rem;
`;

const rightBlockHeader = css`
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 1rem;

  h1 {
    margin: 0;
  }
  svg {
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
