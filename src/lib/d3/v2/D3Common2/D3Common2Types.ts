import { Selection } from 'd3';

/**
 * Types that have dependencies on types defined in d3
 */
export type D3SelectionSVG = Selection<SVGSVGElement, unknown, null, undefined>;
export type D3Selection = Selection<SVGGElement, unknown, null, undefined>;

/**
 * Types that do not depend on the types defined in d3
 */
export type D3Margin = {
  left: number;
  right: number;
  top: number;
  bottom: number;
};
export type D3DoubleNumberArray = [number, number];
export type D3Data = Record<string, any>;
export type D3ScaleType = 'number' | 'time';

/**
 * function params
 */
export type AppendSVGParams = {
  container: HTMLElement;
  width: number;
  height: number;
  className?: string;
};
