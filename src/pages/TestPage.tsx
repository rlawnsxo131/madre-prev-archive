import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import { D3LineChart } from '../lib/d3';
import { D3Data } from '../lib/d3/d3Common/d3CommonTypes';
import { getRandomIntInclusive } from '../lib/utils';

interface TestPageProps {}

function TestPage(props: TestPageProps) {
  const ref = useRef<HTMLDivElement | null>(null);
  const min = 50;
  const max = 87;
  const data: D3Data = Array.from({ length: 51 }).map((_, i) => [
    i * 2,
    getRandomIntInclusive(min, max),
  ]);

  const width = 460;
  const height = 400;
  const fontSize = 10;
  const fontSizeThreeTimes = fontSize * 3;
  const ticks = 30;
  const tickSize = 6;

  useEffect(() => {
    if (!ref.current) return;

    const chart = new D3LineChart({
      container: ref.current,
      width: width,
      height: height,
      xDomain: [0, 100],
      yDomain: [0, 100],
      xRange: [0, width - (fontSizeThreeTimes + ticks + tickSize)],
      yRange: [height - (fontSizeThreeTimes + ticks + tickSize), 0],
      data,
    });
    chart.setAxis({
      xTicks: ticks,
      yTicks: ticks,
      xTickSize: tickSize,
      yTickSize: tickSize,
      xClass: 'x-class',
      yClass: 'y-class',
      axisFontSize: fontSize,
    });
    chart.setLine({});
  }, [ref.current]);

  return <div css={block(fontSize)} ref={ref}></div>;
}

const block = (fontSize: number) => css`
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  .x-class,
  .y-class {
    font-size: ${fontSize}px;
  }
`;

export default TestPage;
