module.exports = {
  preset: 'ts-jest',
  transform: {
    '^.+\\.(js|jsx)?$': 'babel-jest',
    '^.+\\.(ts|tsx)?$': 'ts-jest',
  },
  testEnvironment: 'jsdom',
  moduleFileExtensions: ['js', 'json', 'jsx', 'ts', 'tsx', 'json'],
  testMatch: ['<rootDir>/src/**/**/*.test.(ts|tsx)'],
  transformIgnorePatterns: ['<rootDir>/node_modules/'],
  //   setupFiles: ['dotenv/config'],
};
