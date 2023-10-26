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
  fullWidth as _fullWidth,
  radius as _radius,
  size as _size,
  theme as _theme,
} from './Button.css';

export type ButtonProps<E extends ElementType> = ComponentPropsWithoutRef<E> & {
  as?: E;
  children?: ReactNode;
  variant?: keyof typeof _theme;
  theme?: keyof (typeof _theme)['solid' | 'outline'];
  radius?: keyof typeof _radius;
  size?: keyof typeof _size;
  fullWidth?: boolean;
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
    fullWidth = false,
    ...props
  }: ButtonProps<E>,
  ref?: ComponentPropsWithRef<E>['ref'],
) {
  const Element = as || 'button';

  return (
    <Element
      ref={ref}
      className={classNames(
        block,
        _size[size],
        _theme[variant][theme],
        _radius[radius],
        { [_fullWidth]: fullWidth },
        className,
      )}
      {...props}
    >
      {children}
    </Element>
  );
});
