import {
  AppendSvgParams,
  D3Data,
  D3DoubleNumberArray,
  D3TickFormat,
} from '../D3Common/D3CommonTypes';

export type D3AxisChartConstructorParams = {
  xDomain: number[];
  yDomain: number[];
  xRange: D3DoubleNumberArray;
  yRange: D3DoubleNumberArray;
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

export type D3AxisChartSetAxisBackgroundGridParams = Partial<{
  xClass: string;
  yClass: string;
  xTickFormat: D3TickFormat;
  yTickFormat: D3TickFormat;
}>;

export type D3AxisChartLineType = 'STRAIGHT' | 'CURVE';

export type D3AxisChartSetLineParams = {
  data: D3Data;
  color?: string;
  strokeWidth?: number;
  lineType?: D3AxisChartLineType;
  animate?: boolean;
};
