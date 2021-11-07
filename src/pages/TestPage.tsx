import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import { D3AxisChart } from '../lib/d3';
import { D3Data } from '../lib/d3/D3Common/D3CommonTypes';
import {
  generateUUID,
  getRandomColors,
  getRandomIntInclusive,
} from '../lib/utils';
import { palette } from '../styles';

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
      margin,
    });

    chart.setAxis({
      xTicks: 2,
      yTicks: 5,
      xTickSize: 0,
      yTickSize: 0,
      xClass: 'x-axis',
      yClass: 'y-axis',
      axisFontSize: 12,
    });

    chart.setAxisBackgroundGrid({
      direction: {
        x: true,
        y: false,
      },
      xClass: 'x-grid',
      yClass: 'y-grid',
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
        strokeWidth: 2,
        // lineType: 'CURVE',
        animate: true,
        uuid,
        linejoinType: 'round',
        linecapType: 'round',
        isMouseOverAction: true,
      });
      chart.drawArea({
        data: v,
        color: colors[i],
        animate: true,
        uuid,
        areaType: 'boundary',
        isMouseOverAction: true,
        mouseOverOpacity: 0.8,
        // opacity: 0.3,
      });
      chart.drawCircle({
        data: v,
        color: colors[i],
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

  .x-axis {
    & path {
      color: ${palette.gray['200']};
    }
  }

  .y-axis {
    & path {
      display: none;
    }
  }

  .y-gird {
    color: ${palette.gray['200']};
  }
  .x-grid {
    color: ${palette.gray['200']};
    & path {
      display: none;
    }
  }
`;

export default TestPage;
