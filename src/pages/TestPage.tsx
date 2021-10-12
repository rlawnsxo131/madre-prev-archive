import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import { D3AxisChart } from '../lib/d3';
import { D3Data } from '../lib/d3/D3Common/D3CommonTypes';
import { getRandomColors, getRandomIntInclusive } from '../lib/utils';
import palette from '../styles/palette';

interface TestPageProps {}

function TestPage(props: TestPageProps) {
  const ref = useRef<HTMLDivElement>(null);
  const dataList: D3Data[] = Array.from({ length: 5 }).map((_, i) => ({
    d3Position: Array.from({ length: 11 }).map((v, j) => [
      j * 10,
      getRandomIntInclusive(10, 87),
    ]),
    key: i,
  }));

  const width = 460;
  const height = 400;
  const fontSize = 10;
  const axisMaxUnitExpressionLength = fontSize * 3;
  const xTicks = 5;
  const yTicks = 10;
  const tickSize = 6;
  const xDomain = [0, 100];
  const yDomain = [0, 100];
  const strokeWidth = 2;

  useEffect(() => {
    if (!ref.current) return;
    const chart = new D3AxisChart({
      container: ref.current,
      width: width,
      height: height,
      className: 'axis-chart',
      xDomain,
      yDomain,
      xRange: [0, width - (axisMaxUnitExpressionLength + xTicks + tickSize)],
      yRange: [height - (axisMaxUnitExpressionLength + yTicks + tickSize), 0],
    });

    chart.setAxis({
      xTicks,
      yTicks,
      xTickSize: tickSize,
      yTickSize: tickSize,
      xClass: 'x-axis',
      yClass: 'y-axis',
      axisFontSize: fontSize,
      axisMaxUnitExpressionLength,
    });

    chart.setAxisBackgroundGrid({
      xClass: 'grid',
      yClass: 'grid',
    });

    chart.drawAxis();
    chart.drawGrid();

    const colors = getRandomColors(dataList.length);

    dataList.forEach((v, i) => {
      chart.drawLine({
        data: v,
        color: colors[i],
        strokeWidth,
        lineType: 'CURVE',
        animate: true,
      });
    });
  }, [ref.current]);

  useEffect(() => {
    if (!ref.current) return;
    const chart = new D3AxisChart({
      container: ref.current,
      width: width,
      height: height,
      className: 'axis-chart',
      xDomain,
      yDomain,
      xRange: [0, width - (axisMaxUnitExpressionLength + xTicks + tickSize)],
      yRange: [height - (axisMaxUnitExpressionLength + yTicks + tickSize), 0],
    });

    chart.setAxis({
      xTicks,
      yTicks,
      xTickSize: tickSize,
      yTickSize: tickSize,
      xClass: 'x-axis',
      yClass: 'y-axis',
      axisFontSize: fontSize,
      axisMaxUnitExpressionLength,
    });

    chart.setAxisBackgroundGrid({
      xClass: 'grid',
      yClass: 'grid',
    });

    chart.drawAxis();
    chart.drawGrid();

    const colors = getRandomColors(dataList.length);

    dataList.forEach((v, i) => {
      chart.drawLine({
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

  .x-axis,
  .y-axis {
    & path {
      color: ${palette.gray['200']};
    }
  }
  .grid {
    color: ${palette.gray['200']};
  }
`;

export default TestPage;
