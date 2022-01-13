import { v4 } from 'uuid';
import { getRandomColors } from '../../utils';

export type D3UniqIdentifierValueAndColorArray = {
  key: string;
  color: string;
}[];

export function generateUniqIdentifierValueAndColorArray(
  length: number,
): D3UniqIdentifierValueAndColorArray {
  const colors = getRandomColors(length);
  return Array.from({ length }).map((v, i) => ({
    key: v4(),
    color: colors[i],
  }));
}
