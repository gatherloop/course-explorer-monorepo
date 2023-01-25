import { TamaguiProvider } from 'tamagui';
import { appConfig } from '../src/tamagui.config';

export const decorators = [
  (Story) => <TamaguiProvider config={appConfig}>{Story()}</TamaguiProvider>,
];
