import { PrepareAssetsExecutorSchema } from './schema';
import executor from './executor';

const options: PrepareAssetsExecutorSchema = {};

describe('PrepareAssets Executor', () => {
  it('can run', async () => {
    const output = await executor(options);
    expect(output.success).toBe(true);
  });
});