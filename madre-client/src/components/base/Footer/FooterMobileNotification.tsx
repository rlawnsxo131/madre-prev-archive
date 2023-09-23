import { css } from '@emotion/react';
import { NotificationIcon } from '../../../image/icons';

interface FooterMobileNotificationProps {}

function FooterMobileNotification(props: FooterMobileNotificationProps) {
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
`;

export default FooterMobileNotification;
