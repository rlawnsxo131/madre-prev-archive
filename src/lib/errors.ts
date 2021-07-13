export function throwNotFoundTopic(topic: string) {
  throw new Error(`not found ${topic}`);
}
