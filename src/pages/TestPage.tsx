import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import { D3AxisChart } from '../lib/d3';
import { D3Data } from '../lib/d3/D3Common/D3CommonTypes';
import {
  generateUUID,
  getRandomColors,
  getRandomIntInclusive,
} from '../lib/utils';
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
  const margin = {
    left: 30,
    right: 30,
    top: 30,
    bottom: 30,
  };
  const fontSize = 10;
  const xTicks = 2;
  const yTicks = 5;
  const tickSize = 0;
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
      margin,
    });

    chart.setAxis({
      xTicks,
      yTicks,
      xTickSize: tickSize,
      yTickSize: tickSize,
      xClass: 'x-axis',
      yClass: 'y-axis',
      axisFontSize: fontSize,
    });

    chart.setAxisBackgroundGrid({
      direction: {
        x: true,
        y: false,
      },
      xClass: 'grid',
      yClass: 'grid',
      xTicks: 5,
      // yTicks: 5,
    });

    chart.drawAxis();
    chart.drawGrid();

    const colors = getRandomColors(dataList.length);

    dataList.forEach((v, i) => {
      const uuid = generateUUID();
      chart.drawLine({
        data: v,
        color: colors[i],
        strokeWidth,
        lineType: 'CURVE',
        animate: true,
        uuid,
      });
      chart.drawArea({
        data: v,
        color: colors[i],
        animate: true,
        uuid,
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
