import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import { D3AxisChart } from '../lib/d3';
import { D3Data } from '../lib/d3/d3Common/d3CommonTypes';
import { getRandomColors, getRandomIntInclusive } from '../lib/utils';
import palette from '../styles/palette';

interface TestPageProps {}

function TestPage(props: TestPageProps) {
  const ref = useRef<HTMLDivElement>(null);
  const min = 10;
  const max = 87;
  const dataList: D3Data[] = Array.from({ length: 5 }).map(() => {
    const data: D3Data = Array.from({ length: 11 }).map((_, i) => [
      i * 10,
      getRandomIntInclusive(min, max),
    ]);
    return data;
  });

  const strokeWidth = 1;
  const width = 460;
  const height = 400;
  const fontSize = 10;
  const axisMaxUnitExpressionLength = fontSize * 3;
  const ticks = 10;
  const tickSize = 6;
  const xDomain = [0, 100];
  const yDomain = [0, 100];

  useEffect(() => {
    if (!ref.current) return;
    const chart = new D3AxisChart({
      container: ref.current,
      width: width,
      height: height,
      className: 'axis-chart',
      xDomain,
      yDomain,
      xRange: [0, width - (axisMaxUnitExpressionLength + ticks + tickSize)],
      yRange: [height - (axisMaxUnitExpressionLength + ticks + tickSize), 0],
    });

    chart.setAxis({
      xTicks: 6,
      yTicks: 10,
      xTickSize: tickSize,
      yTickSize: tickSize,
      xClass: 'x-class',
      yClass: 'y-class',
      axisMaxUnitExpressionLength,
      xGridClass: 'grid',
      yGridClass: 'grid',
    });

    const colors = getRandomColors(dataList.length);

    dataList.forEach((v, i) => {
      chart.setLine({
        data: v,
        color: colors[i],
        strokeWidth,
        lineType: 'CURVE',
        animate: true,
      });
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

  .grid {
    color: ${palette.gray[300]};
  }
`;

export default TestPage;
