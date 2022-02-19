import { css } from '@emotion/react';
import usePopupLoginActions from '../../../hooks/core/usePopupLoginActions';
import usePopupLoginState from '../../../hooks/core/usePopupLoginState';
import CloseIcon from '../../../image/icons/CloseIcon';
import { MobileUserImage } from '../../../image/images';
import { mediaQuery, palette } from '../../../styles';
import GoogleAuthButton from '../../auth/GoogleAuthButton';
import PopupBase from '../PopupBase';

interface PopupLoginProps {}

function PopupLogin(props: PopupLoginProps) {
  const { visible } = usePopupLoginState();
  const { onClose } = usePopupLoginActions();

  return (
    <PopupBase visible={visible}>
      <div css={block}>
        <div css={leftBlock}>
          <MobileUserImage />
        </div>
        <div css={rightBlock}>
          <div css={rightBlockHeader}>
            <CloseIcon onClick={onClose} />
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
  justify-content: flex-end;
  align-items: center;
  svg {
    width: 1.25rem;
    height: 1.25rem;
    color: ${palette.gray['600']};
    cursor: pointer;
  }
`;

export default PopupLogin;
