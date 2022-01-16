import { css } from '@emotion/react';
import { useState } from 'react';
import { DropArrowIcon, UserIcon } from '../../../image/icons';
import { palette } from '../../../styles';
import Button from '../../common/Button';

interface UserPersonalMenuProps {}

function UserPersonalMenu(props: UserPersonalMenuProps) {
  const [temp, setTemp] = useState(false);
  return (
    <div css={block}>
      {!temp && (
        <Button
          shape="round"
          color="pink"
          onClick={() => setTemp((prev) => !prev)}
        >
          로그인
        </Button>
      )}
      {temp && (
        <>
          <UserIcon />
          <DropArrowIcon />
        </>
      )}
    </div>
  );
}

const block = css`
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  svg {
    fill: ${palette.gray['900']};
    &:nth-of-type(1) {
      width: 1.875rem;
      height: 1.875rem;
    }
    &:nth-of-type(2) {
      width: 0.5rem;
      height: 0.5rem;
    }
  }
  &:hover {
    svg {
      fill: ${palette.gray['600']};
    }
  }
`;

export default UserPersonalMenu;
