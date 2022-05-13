import { css } from '@emotion/react';
import useScreenSignUpState from '../../../hooks/screenSignUp/useScreenSignUpState';
import { mediaQuery, themePalette } from '../../../styles';
import { SnowBackground } from '../Background';
import ScreenBase from '../ScreenBase';
import ScreenSignUpLeftBlock from './ScreenSignUpLeftBlock';
import ScreenSignUpRightBlockForm from './ScreenSignUpRightBlockForm';

interface ScreenSignUpProps {}

function ScreenSignUp(props: ScreenSignUpProps) {
  const { visible } = useScreenSignUpState();

  return (
    <ScreenBase visible={visible}>
      <SnowBackground withLogo />
      <div css={block}>
        <div css={content}>
          <ScreenSignUpLeftBlock />
          <ScreenSignUpRightBlockForm />
        </div>
      </div>
    </ScreenBase>
  );
}

const block = css`
  position: relative;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
`;

const content = css`
  display: flex;
  background: ${themePalette.bg_element1};
  border: 1px solid ${themePalette.border_element1};
  border-radius: 0.25rem;
  ${mediaQuery(512)} {
    width: 30rem;
  }
`;

export default ScreenSignUp;
