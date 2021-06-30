import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import Three from '../lib/three/Three';
import palette from '../styles/palette';

interface ThreeTestPageProps {}

function ThreeTestPage(props: ThreeTestPageProps) {
  const rendererRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    if (!rendererRef.current) return;
    const THREE = new Three({
      cameraOption: {
        fov: 75,
        aspect: window.innerWidth / window.innerHeight,
        near: 1,
        far: 1000,
      },
      rendererOption: {
        width: window.innerWidth,
        height: window.innerHeight,
      },
    });
    const scene = THREE.getScene();
    const camera = THREE.getCamera();
    const renderer = THREE.getRenderer();
    rendererRef.current.appendChild(renderer.domElement);
    THREE.setCube({ color: palette.red['500'] });
    function animate() {
      requestAnimationFrame(animate);
      const cube = THREE.getCube();
      if (!cube) return;
      cube.rotation.x += 0.01;
      cube.rotation.y += 0.01;
      renderer.render(scene, camera);
    }
    animate();
  }, [rendererRef.current]);
  return <div css={block} ref={rendererRef}></div>;
}

const block = css``;

export default ThreeTestPage;
