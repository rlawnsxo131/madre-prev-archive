export function normalizeString(text: string) {
  return text.replace(/(\s*)/gi, '');
}

export function googlePhotoUrlSizeChange(url: string) {
  return `${url.split('=')[0]}=s300`;
}
