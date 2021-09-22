import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import apiClient from '../api/apiClient';
import { D3AxisChart } from '../lib/d3';
import { D3Data } from '../lib/d3/d3Common/d3CommonTypes';
import { getRandomIntInclusive } from '../lib/utils';
import palette from '../styles/palette';

interface TestPageProps {}

function TestPage(props: TestPageProps) {
  const ref = useRef<HTMLDivElement | null>(null);
  const min = 50;
  const max = 87;
  const data: D3Data = Array.from({ length: 401 }).map((_, i) => [
    i * 0.25,
    getRandomIntInclusive(min, max),
  ]);

  const strokeWidth = 1;
  const width = 460;
  const height = 400;
  const fontSize = 10;
  const axisMaxUnitExpressionLength = fontSize * 3;
  const ticks = 25;
  const tickSize = 6;
  const xDomain = [0, 100];
  const yDomain = [0, 100];

  useEffect(() => {
    if (!ref.current) return;
    const chart = new D3AxisChart({
      container: ref.current,
      width: width,
      height: height,
      xDomain,
      yDomain,
      xRange: [0, width - (axisMaxUnitExpressionLength + ticks + tickSize)],
      yRange: [height - (axisMaxUnitExpressionLength + ticks + tickSize), 0],
      data,
    });

    chart.setAxis({
      xTicks: ticks,
      yTicks: ticks,
      xTickSize: tickSize,
      yTickSize: tickSize,
      xClass: 'x-class',
      yClass: 'y-class',
      axisMaxUnitExpressionLength,
    });

    chart.setLine({
      color: palette.teal['700'],
      strokeWidth,
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
