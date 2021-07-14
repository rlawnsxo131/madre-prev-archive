import { D3DoubleNumberArray } from './d3CommonTypes';

export type D3LineChartConstructorParams = {
  container: HTMLElement;
  width: number;
  height: number;
  className?: string;
  xDomainData: number[];
  yDomainData: number[];
  xRange: D3DoubleNumberArray;
  yRange: D3DoubleNumberArray;
};
