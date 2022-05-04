import { css } from '@emotion/react';
import { themePalette } from '../../../styles';

interface UserHeaderMenuAuthButtonProps {
  show: () => void;
}

function UserHeaderMenuAuthButton({ show }: UserHeaderMenuAuthButtonProps) {
  return (
    <button css={block} onClick={show}>
      로그인
    </button>
  );
}

const block = css`
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

export default UserHeaderMenuAuthButton;
