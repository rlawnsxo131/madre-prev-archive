import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import { D3LineChart } from '../lib/d3';
import { D3Data } from '../lib/d3/d3Common/d3CommonTypes';
import { getRandomIntInclusive } from '../lib/utils';

interface TestPageProps {}

function TestPage(props: TestPageProps) {
  const ref = useRef<HTMLDivElement | null>(null);
  const min = 40;
  const max = 87;
  const data: D3Data = Array.from({ length: 50 }).map((_, i) => [
    i * 2,
    getRandomIntInclusive(min, max),
  ]);
  const xData = data.map((v) => v[0]);
  const yData = data.map((v) => v[1]);

  useEffect(() => {
    if (!ref.current) return;

    const width = 460;
    const height = 400;

    const chart = new D3LineChart({
      container: ref.current,
      width: width,
      height: height,
      xDomain: [0, 100],
      yDomain: [0, 100],
      xRange: [0, width - 24 - 16],
      yRange: [height - 24 - 16, 0],
    });
    chart.setAxis({
      xTicks: 10,
      yTicks: 10,
      xTickSize: 3,
      yTickSize: 3,
      xClassName: 'x-class',
      yClassName: 'y-class',
    });
  }, [ref.current]);

  return <div css={block} ref={ref}></div>;
}

const block = css`
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
`;

export default TestPage;
