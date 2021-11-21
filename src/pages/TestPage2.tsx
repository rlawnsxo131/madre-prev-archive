import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import { D3AxisChart2 } from '../lib/d3';

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

    const chart = new D3AxisChart2({
      container: ref.current,
      width,
      height,
      margin,
    });

    // chart.setData();
    // chart.setDomain();
    // chart.setAxisOptions({});
    // chart.setAxisSvg();
  }, [ref.current]);

  return <div css={block} ref={ref}></div>;
}

const block = css``;

export default TestPage2;
