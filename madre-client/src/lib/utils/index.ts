export function normalizeString(text: string) {
  return text.replace(/(\s*)/gi, '');
}

export function isNormalEnglishString(text: string) {
  return /^[a-zA-Z0-9]{1,}$/.test(text);
}

export function googlePhotoUrlSizeChange(url: string) {
  return `${url.split('=')[0]}=s300`;
}
