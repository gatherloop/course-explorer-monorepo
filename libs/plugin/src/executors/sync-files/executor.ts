import { ExecutorContext } from '@nrwl/devkit';
import * as fs from 'fs-extra';
import { resolve } from 'path';
import { SyncFilesExecutorSchema } from './schema';

export default async function runExecutor(
  options: SyncFilesExecutorSchema,
  context: ExecutorContext
) {
  console.log('Executor ran for SyncFiles', options);

  const sourceFolderAbsolutePath = resolve(context.root, options.sourceFolder);
  const targetFolderAbsolutePath = resolve(context.root, options.targetFolder);

  if (!fs.existsSync(sourceFolderAbsolutePath)) {
    throw new Error(`Source folder not found: ${sourceFolderAbsolutePath}`);
  }

  if (!fs.existsSync(targetFolderAbsolutePath)) {
    fs.mkdirpSync(targetFolderAbsolutePath);
  }

  fs.copySync(sourceFolderAbsolutePath, targetFolderAbsolutePath);

  return {
    success: true,
  };
}
