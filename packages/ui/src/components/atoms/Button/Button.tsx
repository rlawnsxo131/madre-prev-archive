import type { ButtonHTMLAttributes } from 'react';

import { vars } from '@/styles';

import { block } from './Button.css';

type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  size: 'small' | 'medium' | 'large' | 'responsive';
};

export function Button({ children, ...props }: ButtonProps) {
  console.log(vars);
  return (
    <button className={block} {...props}>
      {children}
    </button>
  );
}
