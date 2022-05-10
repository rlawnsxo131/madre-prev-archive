import { memo } from 'react';
import { css } from '@emotion/react';
import { UserIcon } from '../../../image/icons';
import { googlePhotoUrlSizeChange } from '../../../lib/utils';

interface HeaderUserMenuIconProps {
  onClick: () => void;
  photo_url?: string;
}

function HeaderUserMenuIcon({ onClick, photo_url }: HeaderUserMenuIconProps) {
  return (
    <div css={block} onClick={onClick}>
      {photo_url ? (
        <img src={googlePhotoUrlSizeChange(photo_url)} />
      ) : (
        <UserIcon />
      )}
    </div>
  );
}

const block = css`
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0.25rem;
  img {
    width: 1.9rem;
    height: 1.9rem;
    border-radius: 100%;
  }
`;

export default memo(HeaderUserMenuIcon);
