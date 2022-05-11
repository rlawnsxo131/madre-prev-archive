import { memo, useMemo } from 'react';
import { css } from '@emotion/react';
import { UserIcon } from '../../../image/icons';
import { googlePhotoUrlSizeChange } from '../../../lib/utils';
import { basicStyles } from '../../../styles';

interface HeaderUserMenuIconProps {
  onClick: () => void;
  photo_url?: string;
}

function HeaderUserMenuIcon({ onClick, photo_url }: HeaderUserMenuIconProps) {
  const resizePhotoUrl = useMemo(() => {
    if (!photo_url) return '';
    return googlePhotoUrlSizeChange(photo_url);
  }, [photo_url]);

  return (
    <button css={[basicStyles.button, button]} onClick={onClick}>
      {photo_url ? <img src={resizePhotoUrl} /> : <UserIcon />}
    </button>
  );
}

const button = css`
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
