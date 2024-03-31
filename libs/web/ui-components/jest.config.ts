/* eslint-disable */
export default {
  displayName: 'ui-components',
  preset: '../../../jest.preset.js',
  transform: {
    '^.+\\.[tj]sx?$': ['ts-jest', { tsconfig: '<rootDir>/tsconfig.spec.json' }],
  },
  moduleFileExtensions: ['ts', 'tsx', 'js', 'jsx'],
  coverageDirectory: '../../../coverage/libs/web/ui-components',
  setupFilesAfterEnv: ['<rootDir>/spec/setup.ts'],
};
