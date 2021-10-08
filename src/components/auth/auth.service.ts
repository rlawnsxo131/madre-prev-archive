class AuthService {
  private static instance: AuthService;

  private constructor() {}

  static getInstance() {
    if (!this.instance) {
      this.instance = new AuthService();
    }
    return this.instance;
  }
}

export default AuthService.getInstance();
