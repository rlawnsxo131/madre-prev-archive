import usePopupAuthActions from '../../../hooks/core/usePopupAuthActions';
import Button from '../../common/Button';

interface UserPersonalMenuAuthButtonProps {}

function UserPersonalMenuAuthButton(props: UserPersonalMenuAuthButtonProps) {
  const { onShow } = usePopupAuthActions();

  return (
    <Button shape="round" color="pink" onClick={onShow}>
      로그인
    </Button>
  );
}

export default UserPersonalMenuAuthButton;
