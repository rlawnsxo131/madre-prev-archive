import { AppendSVGParams, D3Margin } from '../D3Common2/D3Common2Types';

export type D3AxisChartConstructorParams = {
  width: number;
  height: number;
  margin: D3Margin;
} & AppendSVGParams;

export type D3AxisChartSetAxisOptionsParams = Partial<{
  axisXTicks: number;
  axisYTicks: number;
  axisXTickSize: number;
  axisYTickSize: number;
  axisXClass: string;
  axisYClass: string;
  axisFontSize: number;
}>;
