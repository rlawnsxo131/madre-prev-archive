import { css } from '@emotion/react';
import usePopupAuthActions from '../../../hooks/popupAuth/usePopupAuthActions';
import usePopupAuthState from '../../../hooks/popupAuth/usePopupAuthState';
import { mediaQuery } from '../../../styles';
import { themePalette } from '../../../styles';
import PopupBase from '../PopupBase';
import PopupAuthStyles from './PopupAuth.styles';
import PopupAuthLeftBlock from './PopupAuthLeftBlock';
import PopupAuthRightBlock from './PopupAuthRightBlock';

interface PopupAuthProps {}

function PopupAuth(props: PopupAuthProps) {
  const { visible, isError } = usePopupAuthState();
  const { close, resetError } = usePopupAuthActions();

  return (
    <PopupBase visible={visible}>
      <div css={block(isError)}>
        <PopupAuthLeftBlock />
        <PopupAuthRightBlock
          isError={isError}
          close={close}
          resetError={resetError}
        />
      </div>
    </PopupBase>
  );
}

const block = (isError: boolean) => css`
  display: flex;
  border-radius: 1rem;
  background: ${themePalette.bg_element1};
  ${mediaQuery(512)} {
    width: 30rem;
  }
  ${isError && PopupAuthStyles.shakeAnimation};
`;

export default PopupAuth;
