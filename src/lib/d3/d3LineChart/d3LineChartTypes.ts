import {
  AppendSvgParams,
  D3Data,
  D3DoubleNumberArray,
} from '../d3Common/d3CommonTypes';

export type D3LineChartConstructorParams = {
  xDomain: number[];
  yDomain: number[];
  xRange: D3DoubleNumberArray;
  yRange: D3DoubleNumberArray;
  data: D3Data;
} & AppendSvgParams;

export type D3LineChartSetAxisParams = Partial<{
  xTicks: number;
  yTicks: number;
  xTickSize: number;
  yTickSize: number;
  xClass: string;
  yClass: string;
  axisFontSize: number;
}>;

export type D3LineChartSetLineParams = Partial<{
  color: string;
  strokeWidth: number;
}>;
