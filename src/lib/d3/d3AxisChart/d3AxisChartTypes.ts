import {
  AppendSvgParams,
  D3Data,
  D3DoubleNumberArray,
  D3TickFormat,
} from '../d3Common/d3CommonTypes';

export type D3AxisChartConstructorParams = {
  xDomain: number[];
  yDomain: number[];
  xRange: D3DoubleNumberArray;
  yRange: D3DoubleNumberArray;
  data: D3Data;
} & AppendSvgParams;

export type D3AxisChartSetAxisParams = Partial<{
  xTicks: number;
  yTicks: number;
  xTickSize: number;
  yTickSize: number;
  xClass: string;
  yClass: string;
  axisFontSize: number;
  axisMaxUnitExpressionLength: number;
  xTickFormat: D3TickFormat;
  yTickFormat: D3TickFormat;
}>;

export type D3AxisChartSetLineParams = Partial<{
  color: string;
  strokeWidth: number;
}>;
