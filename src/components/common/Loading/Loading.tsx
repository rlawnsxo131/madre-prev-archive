import { css } from '@emotion/react';
import useLoadingState from '../../../hooks/core/useLoadingState';
import useTransitionTimeoutEffect from '../../../hooks/useTransitionTimeoutEffect';
import { LoadingIcon } from '../../../image/icons';
import { transitions } from '../../../styles';
import OpaqueLayer from '../OpaqueLayer';

interface LoadingProps {}

function Loading(props: LoadingProps) {
  const { visible } = useLoadingState();
  const closed = useTransitionTimeoutEffect({ visible });

  if (!visible && closed) return null;

  return (
    <OpaqueLayer visible={visible}>
      <div css={block}>
        <LoadingIcon />
      </div>
    </OpaqueLayer>
  );
}

const block = css`
  display: flex;
  justify-content: center;
  align-items: center;
  animation: ${transitions.rotation} 1.25s linear infinite;
`;

export default Loading;
