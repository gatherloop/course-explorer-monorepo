import { ExecutorContext } from '@nrwl/devkit';
import { resolve } from 'path';
import * as fs from 'fs-extra';
import { PrepareAssetsExecutorSchema } from './schema';

function matchExtension(filename = '') {
  return function (extension: string) {
    const fileNameSplit = filename.split('.');
    return fileNameSplit[fileNameSplit.length - 1] === extension;
  };
}

function walkDir(
  dir: string,
  whiteList: Array<string>,
  callback: (path: string) => void
) {
  fs.readdirSync(dir).forEach((fileName) => {
    const dirPath = resolve(dir, fileName);
    const isDirectory = fs.statSync(dirPath).isDirectory();

    if (isDirectory) {
      walkDir(dirPath, whiteList, callback);
    } else if (whiteList.find(matchExtension(fileName)) !== undefined) {
      callback(dirPath);
    }
  });
}

export default async function runExecutor(
  options: PrepareAssetsExecutorSchema,
  context: ExecutorContext
) {
  console.log('Executor ran for PrepareAssets', options);

  const assetFolderAbsolutePath = resolve(context.root, options.assetsFolder);

  const pathFiles: Record<string, string | Record<string, string>> = {};

  walkDir(
    assetFolderAbsolutePath,
    ['png', 'ico', 'svg', 'jpg', 'jpeg', 'gif', 'webp'],
    (filePath) => {
      const [assets, module, fileName] = filePath.split('/').slice(-3);
      if (fileName) {
        const [fileNameOnly] = fileName.split('.');
        const assetKey = `${module}/${fileNameOnly}`;
        pathFiles[assetKey] = `/${assets}/${module}/${fileName}`;
      }
    }
  );

  fs.outputFileSync(
    resolve(context.root, options.outputJson),
    JSON.stringify(pathFiles, null, 2)
  );

  return {
    success: true,
  };
}
