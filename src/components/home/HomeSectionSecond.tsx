import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';

interface HomeSectionSecondProps {}

function HomeSectionSecond(props: HomeSectionSecondProps) {
  const canvasRef = useRef<HTMLCanvasElement | null>(null);

  const draw = (ctx: CanvasRenderingContext2D | null) => {
    if (!ctx) return;
    // ctx.beginPath();
    // ctx.fillRect(25, 25, 100, 100);
    // ctx.clearRect(45, 45, 60, 60);
    // ctx.strokeRect(50, 50, 50, 50);
    // ctx.arc(100, 100, 20, 0, Math.PI * 2);
    // ctx.fillStyle = 'black';
    // ctx.fill();
  };

  const animate = () => {
    // requestAnimationFrame();
  };

  useEffect(() => {
    if (!canvasRef?.current) return;
    const ctx = canvasRef.current.getContext('2d');
    draw(ctx);
  }, [canvasRef?.current]);

  return (
    <section css={block}>
      <canvas ref={canvasRef} style={{ border: '1px solid black' }} />
    </section>
  );
}

const block = css`
  display: flex;
  justify-content: center;
`;

export default HomeSectionSecond;
