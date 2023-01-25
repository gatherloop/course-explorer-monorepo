import { AppProps } from 'next/app';
import { TamaguiProvider } from 'tamagui';
import { NextThemeProvider, useRootTheme } from '@tamagui/next-theme';
import { appConfig } from '@course-explorer-monorepo/ui';
import React from 'react';
import '@tamagui/core/reset.css';

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
