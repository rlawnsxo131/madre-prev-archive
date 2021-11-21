import { AppendSVGParams } from '../D3Common2/D3Common2Types';

export type D3AxisChartConstructorParams = {
  width: number;
  height: number;
  margin: D3Margin;
  data: any[];
} & AppendSVGParams;

export type D3Margin = {};
