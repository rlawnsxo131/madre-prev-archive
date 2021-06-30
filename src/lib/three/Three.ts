import * as THREE from 'three';

interface ThreeProps {
  cameraOption: {
    fov: number;
    aspect: number;
    near: number;
    far: number;
  };
  rendererOption: {
    width: number;
    height: number;
  };
}

export default class Three {
  private scene: THREE.Scene;
  private camera: THREE.PerspectiveCamera;
  private renderer: THREE.WebGLRenderer;
  private cube?: THREE.Mesh<THREE.BoxGeometry, THREE.MeshBasicMaterial>;

  constructor({ cameraOption, rendererOption }: ThreeProps) {
    this.scene = new THREE.Scene();
    this.camera = new THREE.PerspectiveCamera(
      cameraOption.fov,
      cameraOption.aspect,
      cameraOption.near,
      cameraOption.far,
    );
    this.renderer = new THREE.WebGLRenderer({
      alpha: true,
    });
    this.renderer.setSize(rendererOption.width, rendererOption.height);
  }

  setCube(meshParams: THREE.MeshBasicMaterialParameters = {}) {
    const geometry = new THREE.BoxGeometry();
    const material = new THREE.MeshBasicMaterial({ ...meshParams });
    this.cube = new THREE.Mesh(geometry, material);
    this.scene.add(this.cube);
    this.camera.position.z = 5;
  }

  getScene() {
    return this.scene;
  }

  getCamera() {
    return this.camera;
  }

  getRenderer() {
    return this.renderer;
  }

  getCube() {
    return this.cube;
  }
}
