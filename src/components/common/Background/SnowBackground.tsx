import { css } from '@emotion/react';
import { palette, zIndexes } from '../../../styles';
import Snowfall from 'react-snowfall';

interface SnowBackgroundProps {
  withLogo?: boolean;
}

function SnowBackground({ withLogo = false }: SnowBackgroundProps) {
  return (
    <div css={block}>
      {withLogo && <h1>Madre</h1>}
      <Snowfall
        snowflakeCount={100}
        color={palette.pink['500']}
        speed={[0.5, 1]}
      />
    </div>
  );
}

const block = css`
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: ${zIndexes.snowBackground};
  h1 {
    position: absolute;
    left: 5%;
  }
`;

export default SnowBackground;
