import { css } from '@emotion/react';
import { format, add } from 'date-fns';
import { useEffect, useRef, useState } from 'react';
import { D3AxisChart2 } from '../lib/d3';
import { getRandomIntInclusive } from '../lib/utils';
import { palette } from '../styles';

interface TestPage2Props {}

function TestPage2(props: TestPage2Props) {
  const ref = useRef<HTMLDivElement>(null);
  const chartRef = useRef<D3AxisChart2 | null>(null);
  const [axisClass, setAxisClass] = useState({
    axisXClass: '',
    axisYClass: '',
  });

  const width = 960;
  const height = 460;
  const margin = {
    left: 100,
    right: 100,
    top: 30,
    bottom: 30,
  };

  const data = Array.from({ length: 5 }).map((_, i) =>
    Array.from({ length: 11 }).map((_, j) => {
      const date = new Date();
      const x = add(date, {
        months: j,
      });
      return {
        x,
        y: getRandomIntInclusive(2000, 10000),
      };
    }),
  );

  useEffect(() => {
    if (!ref.current) return;
    if (chartRef.current) return;

    chartRef.current = new D3AxisChart2({
      container: ref.current,
      width,
      height,
      margin,
    });
  }, [ref.current, chartRef.current]);

  useEffect(() => {
    if (!chartRef.current) return;

    chartRef.current.setData(data);
    chartRef.current.setScaleType('time', 'number');
    chartRef.current.setDomainOptions('x', 'y');
    chartRef.current.setDomain();
    chartRef.current.setAxisOptions({
      axisXTicks: 5,
      axisYTicks: 5,
      axisXTickVisible: true,
      axisYTickVisible: true,
      axisXTickFormat: (d, _) => format(d as Date, 'yyyy-MM'),
      axisFontSize: 12,
    });

    const { axisXClass, axisYClass } = chartRef.current.getAxisClass();
    setAxisClass((prev) => ({
      ...prev,
      axisXClass,
      axisYClass,
    }));

    chartRef.current.setAxis();
    chartRef.current.appendAxis();
    chartRef.current.appendLine();

    setTimeout(() => {
      if (!chartRef.current) return;
      const data = Array.from({ length: 5 }).map((_, i) =>
        Array.from({ length: 11 }).map((_, j) => ({
          x: j * 100,
          y: getRandomIntInclusive(2000, 20000),
        })),
      );
      console.log(data);
      chartRef.current.setData(data);
      chartRef.current.setScaleType('number', 'number');
      chartRef.current.setDomain();
      chartRef.current.setAxisOptions({
        axisXTickFormat: (d, i) => `${d}`,
      });
      chartRef.current.setAxis();
      chartRef.current.updateAxis();
      chartRef.current.updateLine();
    }, 1500);
  }, [chartRef.current]);

  return <div css={block(axisClass)} ref={ref}></div>;
}

const block = ({
  axisXClass,
  axisYClass,
}: {
  axisXClass: string;
  axisYClass: string;
}) => css`
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow-x: scroll;

  .${axisXClass}, .${axisYClass} {
    line,
    path {
      color: ${palette.gray['300']};
    }

    text {
      font-weight: 400;
    }
  }
`;

export default TestPage2;
