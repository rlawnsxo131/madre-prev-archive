import * as d3 from 'd3';
import { throwNotFoundTopic } from '../errors';
import {
  AxisParamType,
  D3NullableSelectionAreaType,
  D3NullableSelectionPathType,
  D3NullableSelectionType,
  DataType,
  D3LineChartSetAreaParams,
  D3LineChartSetCircleParams,
  D3LineChartSetLineParams,
  LineAndPointInitializeParams,
} from './types';

export default class D3LineChart {
  private svg: D3NullableSelectionType = null;
  private data: DataType[] | null = null;
  private xDomain: [number, number] = [0, 0];
  private yDomain: [number, number] = [0, 0];
  private xRange: [number, number] = [0, 0];
  private yRange: [number, number] = [0, 0];
  private path: D3NullableSelectionPathType = null;
  private pathLength?: number = undefined;
  private area: D3NullableSelectionAreaType = null;

  constructor() {}

  initialize({
    svg,
    data,
    xDomainData,
    yDomainData,
    xRange,
    yRange,
    xTransform = '',
    yTransform = '',
    xTickSize,
    yTickSize,
    xTicks,
    yTicks,
  }: LineAndPointInitializeParams) {
    this.svg = d3.select(svg);
    this.data = data;
    this.xDomain = this.getExtent(xDomainData);
    this.yDomain = this.getExtent(yDomainData);
    this.xRange = xRange;
    this.yRange = yRange;
    this.setAxis({
      xTransform,
      yTransform,
      xTickSize,
      yTickSize,
      xTicks,
      yTicks,
    });
  }

  getExtent(data: number[]): [number, number] {
    const [min = 0, max = 0] = d3.extent(data);
    return [min, max];
  }

  xScale() {
    return d3.scaleLinear().domain(this.xDomain).range(this.xRange);
  }

  yScale() {
    return d3.scaleLinear().domain(this.yDomain).range(this.yRange);
  }

  private setAxis({
    xTransform = '',
    yTransform = '',
    xTickSize,
    yTickSize,
    xTicks,
    yTicks,
  }: AxisParamType) {
    if (!this.svg) return throwNotFoundTopic('svg');
    const svg = this.svg;
    const xAxisSvg = svg
      .append('g')
      .attr('font-weight', '500')
      .attr('transform', xTransform);
    const yAxisSvg = svg
      .append('g')
      .attr('font-weight', '500')
      .attr('transform', yTransform);
    const xAxis = d3
      .axisBottom(this.xScale())
      .tickSize(xTickSize)
      .ticks(xTicks);
    const yAxis = d3.axisLeft(this.yScale()).tickSize(yTickSize).ticks(yTicks);
    xAxis(xAxisSvg);
    yAxis(yAxisSvg);
  }

  setCircle({
    xKey,
    yKey,
    color = 'black',
    radius = 3,
  }: D3LineChartSetCircleParams) {
    if (!this.svg) return throwNotFoundTopic('svg');
    if (!this.data) return throwNotFoundTopic('data');
    const xScale = this.xScale();
    const yScale = this.yScale();
    this.svg
      .selectAll('circles')
      .data(this.data)
      .enter()
      .append('circle')
      .attr('fill', color)
      .attr('r', radius)
      .attr('cx', (d) => xScale(d[xKey]))
      .attr('cy', (d) => yScale(d[yKey]));
  }

  setLine({
    xKey,
    yKey,
    color = 'black',
    strokeWidth = 1,
  }: D3LineChartSetLineParams) {
    if (!this.svg) return throwNotFoundTopic('svg');
    if (!this.data) return throwNotFoundTopic('data');
    const xScale = this.xScale();
    const yScale = this.yScale();
    const linearGenerator = d3
      .line<Record<string, any>>()
      .x((d) => xScale(d[xKey]))
      .y((d) => yScale(d[yKey]));
    const path = this.svg
      .append('path')
      .attr('fill', 'none')
      .attr('d', linearGenerator(this.data) as any)
      .attr('stroke-width', strokeWidth)
      .attr('stroke', color);
    this.path = path;
    this.pathLength = path.node()?.getTotalLength();
  }

  setArea({
    xKey,
    yKey,
    color = 'black',
    opacity = 0.3,
  }: D3LineChartSetAreaParams) {
    if (!this.svg) return throwNotFoundTopic('svg');
    if (!this.data) return throwNotFoundTopic('data');
    if (!this.pathLength) return throwNotFoundTopic('pathLength');
    if (!this.yRange[0]) return throwNotFoundTopic('yRange[0]');
    const xScale = this.xScale();
    const yScale = this.yScale();
    const areaGenerator = d3
      .area<DataType>(this.data as any)
      .x((d) => xScale(d[xKey]))
      .y0(this.yRange[0])
      .y1((d) => yScale(d[yKey]));

    const area = this.svg
      .append('path')
      .datum(this.data)
      .attr('fill', color)
      .attr('fill-opacity', opacity)
      .attr('stroke', 'none')
      .attr('d', areaGenerator(this.data) as any);
    this.area = area;
  }

  strokeAnimate(duration = 1500) {
    if (!this.path || !this.pathLength) {
      return throwNotFoundTopic('path and pathLength');
    }
    this.path
      .attr('stroke-dashoffset', this.pathLength)
      .attr('stroke-dasharray', this.pathLength)
      .transition()
      .ease(d3.easeSinInOut)
      .duration(duration)
      .attr('stroke-dashoffset', 0); //시작점
  }

  areaAnimate(duration = 1500, opacity = 0.3) {
    if (!this.area) return throwNotFoundTopic('area');
    this.area
      .attr('fill-opacity', 0)
      .transition()
      .ease(d3.easeSinInOut)
      .duration(duration)
      .attr('fill-opacity', opacity);
  }
}
