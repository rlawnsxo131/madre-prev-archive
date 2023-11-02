import classNames from 'classnames';
import type { PropsWithoutRef } from 'react';
import { forwardRef, type InputHTMLAttributes } from 'react';

import { base, size as _size, warn } from './Input.css';

export type InputProps = PropsWithoutRef<
  Omit<InputHTMLAttributes<HTMLInputElement>, 'size'> & {
    size?: keyof typeof _size;
    error?: boolean;
  }
>;

export const Input = forwardRef<HTMLInputElement, InputProps>(function (
  { className, size = 'medium', error, ...props },
  ref,
) {
  return (
    <input
      ref={ref}
      className={classNames(base, _size[size], { [warn]: error }, className)}
      {...props}
    />
  );
});
