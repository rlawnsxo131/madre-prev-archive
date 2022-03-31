import { css } from '@emotion/react';
import useLoadingState from '../../../hooks/common/useLoadingState';
import useTransitionTimeoutEffect from '../../../hooks/useTransitionTimeoutEffect';
import { LoadingIcon } from '../../../image/icons';
import { transitions, zIndexes } from '../../../styles';
import OpaqueLayer from '../OpaqueLayer';

interface LoadingProps {}

function Loading(props: LoadingProps) {
  const { visible } = useLoadingState();
  const closed = useTransitionTimeoutEffect({ visible });

  if (!visible && closed) return null;

  return (
    <>
      <OpaqueLayer visible={visible} />
      <div css={block}>
        <div css={content}>
          <LoadingIcon />
        </div>
      </div>
    </>
  );
}

const block = css`
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: ${zIndexes.loading};
  display: flex;
  align-items: center;
  justify-content: center;
`;

const content = css`
  display: flex;
  justify-content: center;
  align-items: center;
  animation: ${transitions.rotation} 1.25s linear infinite;
`;

export default Loading;
