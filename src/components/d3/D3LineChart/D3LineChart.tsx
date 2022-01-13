import { format } from 'date-fns';
import { useEffect, useRef } from 'react';
import { v4 } from 'uuid';
import { D3AxisChart, D3FormatUtil, D3GenerateUtil } from '../../../lib/d3';

interface D3LineChartProps<T> {
  width: number;
  height: number;
  margin: {
    left: number;
    right: number;
    top: number;
    bottom: number;
  };
  loading: boolean;
  data: Record<keyof T, any>[][];
}

function D3LineChart<T>({
  width,
  height,
  margin,
  loading,
  data,
}: D3LineChartProps<T>) {
  const containerRef = useRef<HTMLDivElement>(null);
  const chartRef = useRef<D3AxisChart | null>(null);
  const axisXClassRef = useRef<string>(`axis-x-${v4()}`);
  const axisYClassRef = useRef<string>(`axis-x-${v4()}`);

  // initialize chart
  useEffect(() => {
    if (!containerRef.current) return;
    if (chartRef.current) return;
    if (!data) return;
    chartRef.current = new D3AxisChart({
      container: containerRef.current,
      width,
      height,
      margin,
    });
    chartRef.current.setData(data);
    chartRef.current.setUniqIdentifierValueMap(
      D3GenerateUtil.generateUniqIdentifierValueAndColorArray(data.length),
    );
    chartRef.current.setScaleType('time', 'number');
    chartRef.current.setDomainOptions('x', 'y');
    chartRef.current.setDomain();
    chartRef.current.setAxisOptions({
      axisXTicks: 5,
      axisYTicks: 5,
      axisXTickVisible: true,
      axisYTickVisible: true,
      axisXTickFormat: (d, _) => format(d as Date, 'yyyy-MM-dd'),
      axisYTickFormat: (d, _) => D3FormatUtil.formatNumberWithComma()(d),
      axisXClass: axisXClassRef.current,
      axisYClass: axisYClassRef.current,
    });
    chartRef.current.setLineOptions({
      lineStrokeWidth: 2,
      lineType: 'CURVE',
    });
    chartRef.current.setCircleOptions({
      circleDrawDelay: 1000,
    });
    chartRef.current.setAxis();
    chartRef.current.appendAxis();
    chartRef.current.appendLine();
    chartRef.current.appendArea();
    chartRef.current.removeAndAppendCircle();
    chartRef.current.resetData();
  }, [containerRef.current, chartRef.current, data]);

  // update chart
  useEffect(() => {
    if (!containerRef.current) return;
    if (!chartRef.current) return;
    if (loading) return;
    if (!data) return;
    chartRef.current.setData(data);
    chartRef.current.setScaleType('number', 'number');
    chartRef.current.setDomain();
    chartRef.current.setAxisOptions({
      axisXTickFormat: (d, _) => D3FormatUtil.formatNumberWithComma()(d),
    });
    chartRef.current.setAxis();
    chartRef.current.updateAxis();
    chartRef.current.updateLine();
    chartRef.current.updateArea();
    chartRef.current.removeAndAppendCircle();
    chartRef.current.resetData();
  }, [containerRef.current, chartRef.current, loading, data]);

  return <div ref={containerRef}></div>;
}

export default D3LineChart;
