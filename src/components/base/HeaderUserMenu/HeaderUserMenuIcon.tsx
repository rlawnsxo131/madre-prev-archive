import { memo } from 'react';
import { css } from '@emotion/react';
import { UserIcon } from '../../../image/icons';
import { googlePhotoUrlSizeChange } from '../../../lib/utils';
import { themePalette } from '../../../styles';

interface HeaderUserMenuIconProps {
  onClick: () => void;
  photo_url?: string;
}

function HeaderUserMenuIcon({ onClick, photo_url }: HeaderUserMenuIconProps) {
  if (!photo_url) {
    return <UserIcon css={[icon, fill]} onClick={onClick} />;
  }

  return (
    <img
      css={[icon, borderRadius]}
      src={googlePhotoUrlSizeChange(photo_url)}
      onClick={onClick}
    />
  );
}

const icon = css`
  width: 1.9rem;
  height: 1.9rem;
`;

const fill = css`
  fill: ${themePalette.fill1};
`;

const borderRadius = css`
  border-radius: 100%;
`;

export default memo(HeaderUserMenuIcon);
