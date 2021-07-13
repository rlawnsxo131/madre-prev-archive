export type DataType = Record<string, any>;

export type AxisParamType = Pick<
  LineAndPointInitializeParams,
  'xTransform' | 'yTransform' | 'xTickSize' | 'yTickSize' | 'xTicks' | 'yTicks'
>;

export type D3NullableSelectionType = d3.Selection<
  SVGSVGElement,
  unknown,
  null,
  undefined
> | null;

export type D3NullableSelectionPathType = d3.Selection<
  SVGPathElement,
  unknown,
  null,
  undefined
> | null;

export type D3NullableSelectionAreaType = d3.Selection<
  SVGPathElement,
  DataType[],
  null,
  undefined
> | null;

export type D3LineChartKindsType = 'circle' | 'line' | 'area';

export interface D3LineChartSetParams {
  xKey: string;
  yKey: string;
  color?: string;
}

export interface D3LineChartSetCircleParams extends D3LineChartSetParams {
  radius?: number;
}

export interface D3LineChartSetLineParams extends D3LineChartSetParams {
  strokeWidth?: number;
}

export interface D3LineChartSetAreaParams extends D3LineChartSetParams {
  opacity?: number;
}

export interface LineAndPointInitializeParams {
  svg: SVGSVGElement;
  data: DataType[];
  xDomainData: number[];
  yDomainData: number[];
  xRange: [number, number];
  yRange: [number, number];
  xTransform?: string;
  yTransform?: string;
  xTickSize: number;
  yTickSize: number;
  xTicks: number;
  yTicks: number;
}
