import { css } from '@emotion/react';
import { OnlineArtImage } from '../../../image/images';
import { palette } from '../../../styles';

interface ScreenSignupLeftBlockProps {}

function ScreenSignupLeftBlock(props: ScreenSignupLeftBlockProps) {
  return (
    <div css={block}>
      <OnlineArtImage />
    </div>
  );
}

const block = css`
  flex: 1 1 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  background: ${palette.gray['100']};
  border-radius: 0.25rem 0 0 0.25rem;
  svg {
    width: 100%;
  }
`;

export default ScreenSignupLeftBlock;
