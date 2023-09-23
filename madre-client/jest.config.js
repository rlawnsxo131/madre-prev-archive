const esModules = ['d3', 'd3-array', 'other-d3-module-if-needed'].join('|');

module.exports = {
  preset: 'ts-jest',
  transform: {
    '^.+\\.(js|jsx)?$': 'babel-jest',
    '^.+\\.(ts|tsx)?$': 'ts-jest',
  },
  testEnvironment: 'jsdom',
  moduleFileExtensions: ['js', 'json', 'jsx', 'ts', 'tsx', 'json'],
  setupFilesAfterEnv: ['<rootDir>/src/__tests__/config/setupTests.ts'],
  moduleNameMapper: {
    '\\.(css|less)$': '<rootDir>/src/__tests__/mocks/styleMock.ts',
    '^uuid$': require.resolve('uuid'),
  },
  testMatch: ['<rootDir>/src/**/**/*.test.(ts|tsx)'],
  transformIgnorePatterns: [
    '<rootDir>/node_modules/',
    `<rootDir>/node_modules/(?!${esModules})`,
  ],
};
