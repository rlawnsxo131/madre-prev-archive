export function getElement(element: Element | (() => Element) | null) {
  return typeof element === 'function' ? element() : element;
}
