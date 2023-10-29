import classNames from 'classnames';
import type { PropsWithoutRef } from 'react';
import { forwardRef, type InputHTMLAttributes } from 'react';

import { block } from './Input.css';

export type InputProps = PropsWithoutRef<InputHTMLAttributes<HTMLInputElement>>;

export const Input = forwardRef<HTMLInputElement, InputProps>(function (
  { className, ...props },
  ref,
) {
  return (
    <input ref={ref} className={classNames(block, className)} {...props} />
  );
});
