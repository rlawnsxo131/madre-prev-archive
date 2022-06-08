import { css } from '@emotion/react';

interface UserPageMenuProps {}

function UserPageMenu(props: UserPageMenuProps) {
  return <div css={block}>user page menu</div>;
}

const block = css`
  display: flex;
  flex-direction: column;
`;

export default UserPageMenu;
