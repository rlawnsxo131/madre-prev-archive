import { css } from '@emotion/react';
import { format, add } from 'date-fns';
import { useEffect, useRef } from 'react';
import { D3AxisChart2 } from '../lib/d3';
import { getRandomIntInclusive } from '../lib/utils';
import { palette } from '../styles';

interface TestPage2Props {}

function TestPage2(props: TestPage2Props) {
  const ref = useRef<HTMLDivElement>(null);

  const width = 960;
  const height = 460;
  const margin = {
    left: 50,
    right: 50,
    top: 30,
    bottom: 30,
  };

  const data = Array.from({ length: 11 }).map((_, i) => {
    const date = new Date();
    const x = add(date, {
      months: i,
    });
    return {
      x,
      y: getRandomIntInclusive(0, 10000),
    };
  });

  useEffect(() => {
    if (!ref.current) return;

    const chart = new D3AxisChart2({
      container: ref.current,
      width,
      height,
      margin,
    });
    chart.setData(data);
    chart.setScaleType('time', 'number');
    chart.setDomain('x', 'y');
    chart.setAxisOptions({
      axisXTicks: 5,
      axisYTicks: 5,
      axisXTickSize: height,
      axisYTickSize: width,
      axisXTickFormat: (d, i) => format(d as Date, 'yyyy-MM'),
      axisXClass: 'axis-x',
      axisYClass: 'axis-y',
      axisFontSize: 12,
    });
    chart.setAxis();
    chart.appendAxis();

    setTimeout(() => {
      const data = Array.from({ length: 11 }).map((_, i) => ({
        x: i * 100,
        y: getRandomIntInclusive(0, 100),
      }));
      chart.setData(data);
      chart.setScaleType('number', 'number');
      chart.setDomain('x', 'y');
      chart.setAxisTickFormat({
        axisXTickFormat: (d, i) => `${d}`,
      });
      chart.setAxis();
      chart.updateAxis();
    }, 1500);

    setTimeout(() => {
      const data = Array.from({ length: 11 }).map((_, i) => ({
        x: i * 10,
        y: getRandomIntInclusive(0, 100),
      }));
      chart.setData(data);
      chart.setDomain('x', 'y');
      chart.setAxis();
      chart.updateAxis();
    }, 3000);
  }, [ref.current]);

  return <div css={block} ref={ref}></div>;
}

const block = css`
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow-x: scroll;

  .axis-x {
    line,
    path {
      color: ${palette.gray['300']};
    }
  }

  .axis-y {
    line,
    path {
      color: ${palette.gray['300']};
    }
  }
`;

export default TestPage2;
