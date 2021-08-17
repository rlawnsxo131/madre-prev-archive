import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import { D3AxisChart } from '../lib/d3';
import { D3Data } from '../lib/d3/d3Common/d3CommonTypes';
import { getRandomIntInclusive } from '../lib/utils';
import palette from '../styles/palette';

interface TestPageProps {}

function TestPage(props: TestPageProps) {
  const ref = useRef<HTMLDivElement | null>(null);
  const min = 50;
  const max = 87;
  const data: D3Data = Array.from({ length: 101 }).map((_, i) => [
    i * 2,
    getRandomIntInclusive(min, max),
  ]);

  const strokeWidth = 1;
  const width = 460;
  const height = 400;
  const fontSize = 10;
  const maxUnitExpressionLength = fontSize * 3;
  const ticks = 25;
  const tickSize = 6;

  useEffect(() => {
    if (!ref.current) return;

    const chart = new D3AxisChart({
      container: ref.current,
      width: width,
      height: height,
      xDomain: [0, 100],
      yDomain: [0, 100],
      xRange: [0, width - (maxUnitExpressionLength + ticks + tickSize)],
      yRange: [height - (maxUnitExpressionLength + ticks + tickSize), 0],
      data,
    });

    chart.setAxis({
      xTicks: ticks,
      yTicks: ticks,
      xTickSize: tickSize,
      yTickSize: tickSize,
      xClass: 'x-class',
      yClass: 'y-class',
      maxUnitExpressionLength,
    });
    chart.setLine({
      color: palette.teal['700'],
      strokeWidth,
    });
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
