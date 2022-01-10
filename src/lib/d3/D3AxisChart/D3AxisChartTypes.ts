import { curveBasis, curveMonotoneX, curveMonotoneY } from 'd3';
import {
  AppendSVGParams,
  D3Margin,
  D3TickFormat,
} from '../D3Common/D3CommonTypes';

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
  axisXTickVisible: boolean;
  axisYTickVisible: boolean;
  axisXTickFormat: D3TickFormat;
  axisYTickFormat: D3TickFormat;
  axisXClass: string;
  axisYClass: string;
  axisTransitionDuration: number;
}>;

/**
 * line
 */
export type D3AxisChartLineType = 'STRAIGHT' | 'CURVE';

export type D3AxisChartLinecurvKeys =
  | 'curveBasis'
  | 'curveMonotoneX'
  | 'curveMonotoneY';
export type D3AxisChartLinecurvType =
  | typeof curveBasis
  | typeof curveMonotoneX
  | typeof curveMonotoneY;

export type D3AxisChartLinecapType = 'round' | 'butt';
export type D3AxisChartLinejoinType = 'round' | 'miter';

export type D3AxisChartSetLineOptionsParams = Partial<{
  lineType: D3AxisChartLineType;
  linecurvType: D3AxisChartLinecurvKeys;
  linecapType: D3AxisChartLinecapType;
  linejoinType: D3AxisChartLinejoinType;
  lineStrokeWidth: number;
  lineTransitionDuration: number;
}>;

/**
 * area
 */
export type D3AxisChartAreaType = 'full' | 'boundary';
export type D3AxisChartSetAreaOptionsParams = Partial<{
  areaType: D3AxisChartAreaType;
  areaOpacity: number;
  areaTransitionDuration: number;
}>;

/**
 * circle
 */
export type D3AxisChartSetCircleOptionsParams = Partial<{
  circleRadius: number;
  circleStrokeWidth: number;
  circleIsFill: boolean;
  circleDrawDelay: number;
}>;
