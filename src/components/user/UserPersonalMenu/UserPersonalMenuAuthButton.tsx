import { css } from '@emotion/react';
import usePopupLoginActions from '../../../hooks/core/usePopupLoginActions';
import Button from '../../common/Button';

interface UserPersonalMenuAuthButtonProps {}

function UserPersonalMenuAuthButton(props: UserPersonalMenuAuthButtonProps) {
  const { onShow } = usePopupLoginActions();

  return (
    <Button shape="round" color="pink" onClick={onShow}>
      로그인
    </Button>
  );
}

const block = css``;

export default UserPersonalMenuAuthButton;
