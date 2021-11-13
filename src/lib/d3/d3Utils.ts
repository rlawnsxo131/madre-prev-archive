import { extent, randomInt, randomUniform } from 'd3';

function d3Extent() {}

function d3RandomInt(min: number, max: number) {
  return randomInt(min, max)();
}

function d3RandomUniform(min: number, max: number) {
  return randomUniform(min, max)();
}

export default {
  d3Extent,
  d3RandomInt,
  d3RandomUniform,
};
