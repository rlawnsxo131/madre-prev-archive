import { css } from '@emotion/react';
import { memo } from 'react';
import { basicStyles, themePalette } from '../../../styles';

interface HeaderUserMenuAuthButtonProps {
  show: () => void;
}

function HeaderUserMenuAuthButton({ show }: HeaderUserMenuAuthButtonProps) {
  return (
    <button css={[basicStyles.button, block]} onClick={show}>
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
  color: ${themePalette.text1};
  background: none;
`;

export default memo(HeaderUserMenuAuthButton);
