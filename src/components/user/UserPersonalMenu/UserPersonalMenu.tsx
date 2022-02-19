import { css } from '@emotion/react';
import { DropArrowIcon, UserIcon } from '../../../image/icons';
import { palette } from '../../../styles';
import UserPersonalMenuAuthButton from './UserPersonalMenuAuthButton';

interface UserPersonalMenuProps {}

function UserPersonalMenu(props: UserPersonalMenuProps) {
  return (
    <div css={block}>
      <UserPersonalMenuAuthButton />

      {/* <UserIcon />
          <DropArrowIcon />
         */}
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
    fill: ${palette.gray['700']};
    &:nth-of-type(1) {
      width: 1.9rem;
      height: 1.9rem;
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
