import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import { randomColor, randomNumber } from '../../lib/utils';

interface HomeSectionSecondProps {}

function getCircles() {
  const circles = [];
  for (let i = 0; i < 150; i++) {
    const x = randomNumber(15, 315);
    const y = randomNumber(15, 145);
    const size = randomNumber(5, 15);
    const color = randomColor();
    circles.push({
      x,
      y,
      size,
      color,
    });
  }
  return circles;
}

function HomeSectionSecond(props: HomeSectionSecondProps) {
  const circles = getCircles();
  const canvasRef = useRef<HTMLCanvasElement | null>(null);
  const startTime = useRef<number>(0);
  const endTime = useRef<number>(0);
  const operator = useRef<boolean>(false);

  /**
   * change trigonometric functions or three.js ?
   */
  const draw = () => {
    const ctx = canvasRef.current?.getContext('2d');
    if (!ctx) return;
    if (!startTime.current) startTime.current = Date.now();
    ctx.clearRect(0, 0, 320, 150);
    for (let i = 0; i < circles.length; i++) {
      ctx.beginPath();
      ctx.arc(circles[i].x, circles[i].y, circles[i].size, 0, Math.PI * 2);
      ctx.fillStyle = circles[i].color;
      ctx.fill();
      if (operator.current) {
        if (i < circles.length / 2) {
          circles[i].y = circles[i].y - 0.1;
          circles[i].x = circles[i].x - 0.1;
        } else {
          circles[i].y = circles[i].y + 0.1;
          circles[i].x = circles[i].x + 0.1;
        }
      }
      if (!operator.current) {
        if (i < circles.length / 2) {
          circles[i].y = circles[i].y + 0.1;
          circles[i].x = circles[i].x + 0.1;
        } else {
          circles[i].y = circles[i].y - 0.1;
          circles[i].x = circles[i].x - 0.1;
        }
      }
    }
    endTime.current = Date.now();
    if (endTime.current - startTime.current >= 2000) {
      operator.current = !operator.current;
      startTime.current = Date.now();
    }
    window.requestAnimationFrame(draw);
  };

  useEffect(() => {
    if (!canvasRef?.current) return;
    draw();
  }, [canvasRef?.current]);

  return (
    <section css={block}>
      <canvas ref={canvasRef} width={320} height={150} />
    </section>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
  align-items: center;
  canvas {
    background: none;
  }
`;

export default HomeSectionSecond;
