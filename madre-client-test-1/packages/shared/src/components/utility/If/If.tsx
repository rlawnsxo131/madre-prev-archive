import type { PropsWithChildren } from 'react';

export type IfProps = PropsWithChildren<{
  predicate: boolean | (() => boolean);
}>;

export function If({ children, predicate }: IfProps) {
  const condition = typeof predicate === 'function' ? predicate() : predicate;
  return condition ? children : null;
}
