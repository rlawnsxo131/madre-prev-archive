import { v4 } from 'uuid';
import getRandomColors from '../../utils/getRandomColors';
import { D3UniqIdentifierValueAndColorArray } from '../types/D3Util.types';

export function generateUniqIdentifierValueAndColorArray(
  length: number,
): D3UniqIdentifierValueAndColorArray {
  const colors = getRandomColors(length);
  return Array.from({ length }).map((v, i) => ({
    key: v4(),
    color: colors[i],
  }));
}
