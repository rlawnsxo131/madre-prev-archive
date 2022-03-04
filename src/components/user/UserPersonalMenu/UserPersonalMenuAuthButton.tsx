import usePopupAuthActions from '../../../hooks/core/usePopupAuthActions';
import Button from '../../common/Button';

interface UserPersonalMenuAuthButtonProps {}

function UserPersonalMenuAuthButton(props: UserPersonalMenuAuthButtonProps) {
  const { show } = usePopupAuthActions();

  return (
    <Button shape="round" color="pink" onClick={show}>
      로그인
    </Button>
  );
}

export default UserPersonalMenuAuthButton;
