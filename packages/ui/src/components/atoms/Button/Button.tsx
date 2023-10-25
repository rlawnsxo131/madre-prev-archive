import classNames from 'classnames';
import type {
  ComponentPropsWithoutRef,
  ComponentPropsWithRef,
  ElementType,
  ReactNode,
} from 'react';
import { forwardRef } from 'react';

import {
  block,
  radius as _radius,
  size as _size,
  theme as _theme,
} from './Button.css';

export type ButtonProps<E extends ElementType> = ComponentPropsWithoutRef<E> & {
  as?: E;
  children?: ReactNode;
  variant?: 'solid' | 'outline';
  theme?: 'primary' | 'secondary' | 'warn';
  radius?: 'none' | 'medium' | 'full';
  size?: 'small' | 'medium' | 'large' | 'responsive';
};

export type ButtonComponent = <E extends ElementType = 'button'>(
  Props: ButtonProps<E> & { ref?: ComponentPropsWithRef<E>['ref'] },
) => ReactNode;

export const Button: ButtonComponent = forwardRef(function <
  E extends ElementType,
>(
  {
    as,
    children,
    className,
    variant = 'solid',
    theme = 'primary',
    radius = 'medium',
    size = 'medium',
    ...props
  }: ButtonProps<E>,
  ref: ComponentPropsWithRef<E>['ref'],
) {
  const Element = as || 'button';

  return (
    <Element
      ref={ref}
      className={classNames(
        block,
        _theme[variant][theme],
        _size[size],
        _radius[radius],
        className,
      )}
      {...props}
    >
      {children}
    </Element>
  );
});
