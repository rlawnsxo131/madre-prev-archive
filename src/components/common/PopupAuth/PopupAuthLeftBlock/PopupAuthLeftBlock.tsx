import { css } from '@emotion/react';
import { MobileUserImage } from '../../../../image/images';
import { palette } from '../../../../styles';

interface PopupAuthLeftBlockProps {}

function PopupAuthLeftBlock(props: PopupAuthLeftBlockProps) {
  return (
    <div css={block}>
      <MobileUserImage />
    </div>
  );
}

const block = css`
  flex: 1 1 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  background: ${palette.gray['100']};
  border-radius: 0.25rem;
  svg {
    width: 100%;
  }
`;

export default PopupAuthLeftBlock;
