import {
  AppendSvgParams,
  D3DoubleNumberArray,
} from '../d3Common/d3CommonTypes';

export type D3LineChartConstructorParams = {
  xDomain: number[];
  yDomain: number[];
  xRange: D3DoubleNumberArray;
  yRange: D3DoubleNumberArray;
} & AppendSvgParams;

export type D3LineChartSetAxisParams = {
  xTicks?: number;
  yTicks?: number;
  xTickSize?: number;
  yTickSize?: number;
  xClassName?: string;
  yClassName?: string;
};
