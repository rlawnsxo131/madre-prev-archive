import {
  AppendSvgParams,
  D3Data,
  D3TickFormat,
} from '../D3Common/D3CommonTypes';

export type D3Margin = {
  left: number;
  right: number;
  top: number;
  bottom: number;
};

export type D3AxisChartConstructorParams = {
  xDomain: number[];
  yDomain: number[];
  width: number;
  height: number;
  margin: D3Margin;
} & AppendSvgParams;

/**
 * set axis
 */
export type D3AxisChartSetAxisParams = Partial<{
  axisXTicks: number;
  axisYTicks: number;
  axisXTickSize: number;
  axisYTickSize: number;
  axisXClass: string;
  axisYClass: string;
  axisFontSize: number;
}>;

export type D3AxisChartSetAxisBackgroundGridParams = {
  axisBackgroundGridDirection: {
    x: boolean;
    y: boolean;
  };
  axisBackgroundGridXTicks?: number;
  axisBackgroundGridYTicks?: number;
  axisBackgroundGridXTickFormat?: D3TickFormat;
  axisBackgroundGridYTickFormat?: D3TickFormat;
  axisBackgroundGridXClass?: string;
  axisBackgroundGridYClass?: string;
};

/**
 * line
 */
export type D3AxisChartLineType = 'STRAIGHT' | 'CURVE';
export type D3AxisChartCurvType =
  | 'curveBasis'
  | 'curveMonotoneX'
  | 'curveMonotoneY';
export type D3AxisChartLinejoinType = 'round' | 'miter';
export type D3AxisChartLinecapType = 'round' | 'butt';

export type D3AxisChartDrawLineParams = {
  data: D3Data;
  color?: string;
  lineType?: D3AxisChartLineType;
  lineCurvType?: D3AxisChartCurvType;
  lineStrokeWidth?: number;
  linejoinType?: D3AxisChartLinejoinType;
  linecapType?: D3AxisChartLinecapType;
  lineDrawAnimate?: boolean;
  lineDrawAnimateDuration?: number;
  isMouseOverAction?: boolean;
  uuid?: string;
};

/**
 * area
 */
export type AreaType = 'full' | 'boundary';
export type D3AxisChartDrawAreaParams = {
  data: D3Data;
  color?: string;
  areaOpacity?: number;
  areaDrawAnimate?: boolean;
  areaDrawAnimateDuration?: number;
  areaType?: AreaType;
  isMouseOverAction?: boolean;
  areaMouseOverOpacity?: number;
  uuid?: string;
};

/**
 * circle
 */
export type D3AxisChartDrawCircleParams = {
  data: D3Data;
  color?: string;
  circleRadius?: number;
  circleStrokeWidth?: number;
  uuid?: string;
  isMouseOverAction?: boolean;
};
