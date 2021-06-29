export function randomColor() {
  const r = Math.round(Math.random() * 255);
  const g = Math.round(Math.random() * 255);
  const b = Math.round(Math.random() * 255);
  const a = randomNumber(1, 10) * 0.1;
  return `rgba(${r}, ${g}, ${b}, ${a})`;
}

export function randomNumber(min: number, max: number) {
  const randomNumber = Math.random() * (max - min) + min;
  return Math.floor(randomNumber);
}
