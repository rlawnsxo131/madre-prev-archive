import { css } from '@emotion/react';
import { randomColors } from '../../lib/utils';
import d3Data from '../../mocks/d3';
import { LineChart } from '../d3';

interface HomeSectionSecondProps {}

function HomeSectionSecond(props: HomeSectionSecondProps) {
  return (
    <section css={block}>
      <LineChart
        data={d3Data}
        width={400}
        height={330}
        xDomainData={[5, 100]}
        yDomainData={[100, 5]}
        xRange={[25, 390]}
        yRange={[300, 5]}
        xTickSize={5}
        xTicks={15}
        yTickSize={5}
        yTicks={15}
        xTransform={'translate(0, 300)'}
        yTransform={'translate(25,0)'}
        colors={randomColors(4)}
        xKeys={['x', 'x', 'x', 'x']}
        yKeys={['y', 'y1', 'y2', 'y3']}
        kinds={['line', 'area']}
      />
      <LineChart
        data={d3Data}
        width={400}
        height={330}
        xDomainData={[5, 100]}
        yDomainData={[100, 5]}
        xRange={[25, 390]}
        yRange={[300, 5]}
        xTickSize={5}
        xTicks={15}
        yTickSize={5}
        yTicks={15}
        xTransform={'translate(0, 300)'}
        yTransform={'translate(25,0)'}
        colors={randomColors(4)}
        xKeys={['x', 'x', 'x', 'x']}
        yKeys={['y', 'y1', 'y2', 'y3']}
        kinds={['circle', 'line', 'area']}
      />
    </section>
  );
}

const block = css`
  display: flex;
  justify-content: center;
  flex-flow: row wrap;
`;

export default HomeSectionSecond;
