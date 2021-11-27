import {
  AppendSVGParams,
  D3Margin,
  D3TickFormat,
} from '../D3Common2/D3Common2Types';

export type D3AxisChartConstructorParams = {
  width: number;
  height: number;
  margin: D3Margin;
} & AppendSVGParams;

/**
 * axis
 */
export type D3AxisChartSetAxisOptionsParams = Partial<{
  axisXTicks: number;
  axisYTicks: number;
  axisXTickFormat: D3TickFormat;
  axisYTickFormat: D3TickFormat;
  axisXClass: string;
  axisYClass: string;
  axisFontSize: number;
}>;

/**
 * line
 */
export type D3AxisChartLineType = 'STRAIGHT' | 'CURVE';
export type D3AxisChartCurvType =
  | 'curveBasis'
  | 'curveMonotoneX'
  | 'curveMonotoneY';
export type D3AxisChartLinecapType = 'round' | 'butt';
export type D3AxisChartLinejoinType = 'round' | 'miter';

export type D3AxisChartSetLineOptionsParams = Partial<{
  color: string;
  lineType: D3AxisChartLineType;
  lineCurvType: D3AxisChartCurvType;
  lineStrokeWidth: number;
  linecapType: D3AxisChartLinecapType;
  linejoinType: D3AxisChartLinejoinType;
  lineDrawAnimate: boolean;
  lineDrawAnimateDuration: number;
  isMouseOverAction: boolean;
}>;
