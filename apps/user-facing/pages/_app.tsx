import { NextThemeProvider, useRootTheme } from '@tamagui/next-theme';
import { AppProps } from 'next/app';
import React from 'react';
import { TamaguiProvider } from 'tamagui';
import appConfig from '../tamagui.config';

function CustomApp({ Component, pageProps }: AppProps) {
  const [theme, setTheme] = useRootTheme();

  const contents = React.useMemo(
    () => <Component {...pageProps} />,
    [Component, pageProps]
  );

  return (
    <NextThemeProvider onChangeTheme={setTheme}>
      <TamaguiProvider
        disableInjectCSS
        disableRootThemeClass
        defaultTheme={theme}
        config={appConfig}
      >
        {contents}
      </TamaguiProvider>
    </NextThemeProvider>
  );
}

export default CustomApp;
