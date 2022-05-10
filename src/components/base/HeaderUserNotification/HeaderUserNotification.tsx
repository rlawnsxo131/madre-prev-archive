import { css } from '@emotion/react';
import { NotificationIcon } from '../../../image/icons';

interface HeaderUserNotificationProps {}

function HeaderUserNotification(props: HeaderUserNotificationProps) {
  return (
    <div css={block}>
      <NotificationIcon />
    </div>
  );
}

const block = css`
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
`;

export default HeaderUserNotification;
