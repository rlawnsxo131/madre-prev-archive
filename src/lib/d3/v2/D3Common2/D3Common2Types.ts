import { Selection } from 'd3';

/**
 * Types that have dependencies on types defined in d3
 */
export type D3SelectionSVG = Selection<SVGSVGElement, unknown, null, undefined>;

export type AppendSVGParams = {
  container: HTMLElement;
  width: number;
  height: number;
  className?: string;
};
