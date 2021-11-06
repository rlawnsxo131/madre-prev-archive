export default function getPrefersColorScheme() {
  const colorScheme = ['light', 'dark'].reduce((acc, mode) => {
    if (globalThis.matchMedia(`(prefers-color-scheme: ${mode})`).matches) {
      acc += mode;
    }
    return acc;
  }, '');
  return colorScheme || 'light';
}
