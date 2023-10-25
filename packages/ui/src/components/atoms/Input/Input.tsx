import type { PropsWithoutRef } from 'react';
import { forwardRef, type InputHTMLAttributes } from 'react';

export type InputProps = PropsWithoutRef<InputHTMLAttributes<HTMLInputElement>>;

export const Input = forwardRef<HTMLInputElement, InputProps>(function (
  props: InputProps,
  ref,
) {
  return <input ref={ref} {...props} />;
});
