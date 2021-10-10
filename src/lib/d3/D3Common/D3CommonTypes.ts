import { NumberValue, Selection, Axis, Line } from 'd3';

/**
 * Types that have dependencies on types defined in d3
 */
export type D3SelectionSVG = Selection<SVGSVGElement, unknown, null, undefined>;
export type D3Selection = Selection<
  SVGGElement,
  unknown,
  null,
  undefined
> | null;
export type D3Axis = Axis<NumberValue> | null;
export type D3Path = Selection<SVGPathElement, unknown, null, undefined>;
export type D3Line = Line<[number, number]> | null;
export type D3TickFormat = (domainValue: NumberValue, index: number) => string;

/**
 * Types that do not depend on the types defined in d3
 */
export type D3Data = [number, number][] | Iterable<[number, number]>;
export type D3DoubleNumberArray = [number, number];

/**
 * function params
 */
export type AppendSvgParams = {
  container: HTMLElement;
  width: number;
  height: number;
  className?: string;
};
