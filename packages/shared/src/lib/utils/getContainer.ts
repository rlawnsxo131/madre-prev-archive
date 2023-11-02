export function getContainer(container: Element | (() => Element) | null) {
  return typeof container === 'function' ? container() : container;
}
