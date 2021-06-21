import { css } from '@emotion/react';

interface LayoutProps {
  children: React.ReactNode;
}

function Layout({ children }: LayoutProps) {
  return (
    <div css={block}>
      layout
      {children}
    </div>
  );
}

const block = css``;

export default Layout;
