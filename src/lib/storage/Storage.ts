class FallbackStorage {
  private static fallbackStorage: {
    [key: string]: string;
  } = {};

  static valid: boolean = this.checkStorage();

  private static checkStorage() {
    try {
      localStorage.setItem('check', 'check');
      localStorage.removeItem('check');
      return true;
    } catch (e) {
      return false;
    }
  }

  static setItem(key: string, value: any) {
    const string = JSON.stringify(value);
    if (this.valid) {
      localStorage.setItem(key, string);
      return;
    }
    this.fallbackStorage[key] = string;
  }

  static getItem(key: string) {
    const value = this.valid
      ? localStorage.getItem(key)
      : this.fallbackStorage[key];
    try {
      const parsed = JSON.parse(value || '');
      return parsed;
    } catch (e) {
      return null;
    }
  }

  static removeItem(key: string) {
    if (this.valid) {
      localStorage.removeItem(key);
      return;
    }
    delete this.fallbackStorage[key];
  }
}

export default FallbackStorage;
