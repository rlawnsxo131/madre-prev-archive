import { css } from '@emotion/react';
import useScreenSignupState from '../../../hooks/screenSignup/useScreenSignupState';
import Input from '../Input';
import ScreenBase from '../ScreenBase';

interface ScreenSignupProps {}

function ScreenSignup(props: ScreenSignupProps) {
  const { visible } = useScreenSignupState();

  return (
    <ScreenBase visible={visible}>
      <div css={block}>
        <Input />
      </div>
    </ScreenBase>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
`;

export default ScreenSignup;
