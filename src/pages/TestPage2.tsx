import { css } from '@emotion/react';
import { format, add } from 'date-fns';
import { useEffect, useRef } from 'react';
import { D3AxisChart2 } from '../lib/d3';
import { getRandomIntInclusive } from '../lib/utils';

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
      // x: i * 100,
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
    chart.setScaleType('time', 'number');
    chart.setData(data);
    chart.setDomain('x', 'y');
    chart.setAxisOptions({
      axisXTicks: 5,
      axisYTicks: 5,
      axisXTickSize: height,
      axisYTickSize: width,
      axisXTickFormat: (d, i) => format(d as Date, 'yyyy-MM'),
      axisXClass: '',
      axisYClass: '',
      axisFontSize: 12,
    });
    chart.setAxis();
    chart.appendAxis();
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
`;

export default TestPage2;
