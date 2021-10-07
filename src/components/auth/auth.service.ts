class AuthService {
  private static instance: AuthService;

  private constructor() {}

  static getInstance() {
    if (!this.instance) {
      this.instance = new AuthService();
    }
  }
}

export default AuthService.getInstance();
