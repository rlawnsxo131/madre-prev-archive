import { css } from '@emotion/react';
import usePopupAuthActions from '../../../hooks/popupAuth/usePopupAuthActions';
import { themePalette } from '../../../styles';

interface UserMenuButtonAuthProps {}

function UserMenuButtonAuth(props: UserMenuButtonAuthProps) {
  const { show } = usePopupAuthActions();

  return (
    <button css={button} onClick={show}>
      로그인
    </button>
  );
}

const button = css`
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.5rem 0.25rem 0.5rem 0.25rem;
  font-size: 0.9rem;
  font-weight: bold;
  cursor: pointer;
  outline: none;
  border: none;
  box-sizing: border-box;
  cursor: pointer;
  color: ${themePalette.text1};
  background: none;
`;

export default UserMenuButtonAuth;
