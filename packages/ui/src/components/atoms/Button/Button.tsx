import classNames from 'classnames';
import { type ButtonHTMLAttributes } from 'react';

import {
  block,
  radius as _radius,
  size as _size,
  themeOutline,
  themeSolid,
} from './Button.css';

export type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  variant?: 'solid' | 'outline';
  theme?: 'primary' | 'secondary' | 'warn';
  radius?: 'none' | 'medium' | 'full';
  size?: 'small' | 'medium' | 'large' | 'responsive';
};

export function Button({
  children,
  variant = 'solid',
  theme = 'primary',
  radius = 'medium',
  size = 'medium',
  ...props
}: ButtonProps) {
  const _theme = {
    solid: themeSolid,
    outline: themeOutline,
  }[variant][theme];

  return (
    <button
      className={classNames(block, _theme, _size[size], _radius[radius])}
      {...props}
    >
      {children}
    </button>
  );
}
