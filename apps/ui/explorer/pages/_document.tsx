import NextDocument, { Head, Html, Main, NextScript } from 'next/document';
import { Children } from 'react';
import { AppRegistry } from 'react-native';

import { appConfig as Tamagui } from '@course-explorer-monorepo/ui';

export default class Document extends NextDocument {
  static async getInitialProps({ renderPage }) {
    AppRegistry.registerComponent('Main', () => Main);
    const page = await renderPage();
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    const { getStyleElement } = AppRegistry.getApplication('Main');
    const styles = [
      getStyleElement(),
      <style
        key="tamagui-css"
        dangerouslySetInnerHTML={{ __html: Tamagui.getCSS() }}
      />,
    ];
    return { ...page, styles: Children.toArray(styles) };
  }

  render() {
    return (
      <Html>
        <Head></Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}
