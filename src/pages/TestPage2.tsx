import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import { D3AxisChart2 } from '../lib/d3';
import { getRandomIntInclusive } from '../lib/utils';

interface TestPage2Props {}

function TestPage2(props: TestPage2Props) {
  const ref = useRef<HTMLDivElement>(null);

  const width = 460;
  const height = 400;
  const margin = {
    left: 30,
    right: 30,
    top: 30,
    bottom: 30,
  };

  useEffect(() => {
    if (!ref.current) return;

    const data = Array.from({ length: 11 }).map((_, i) => ({
      x: i * 10,
      y: getRandomIntInclusive(10, 87),
    }));

    const chart = new D3AxisChart2({
      container: ref.current,
      width,
      height,
      margin,
    });

    chart.setDataAndDomain({ data, xKey: 'x', yKey: 'y' });
    chart.setAxisOptions({
      axisXTicks: 3,
      axisYTicks: 5,
      axisXTickSize: width,
      axisYTickSize: height,
      axisXClass: '',
      axisYClass: '',
      axisFontSize: 10,
    });
    chart.setAxisSvg();
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

export default TestPage2;
