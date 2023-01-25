import { createTamagui } from 'tamagui';
import { config } from '@tamagui/config-base';

export const appConfig = createTamagui(config);

export type AppConfig = typeof appConfig;

declare module '@tamagui/core' {
  // eslint-disable-next-line @typescript-eslint/no-empty-interface
  interface TamaguiCustomConfig extends AppConfig {}
}

export default appConfig;
