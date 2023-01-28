import assetsMap from '../../__generated__/assets/index.json';

type AssetKey = keyof typeof assetsMap;

export const getAssetSrc = (key: AssetKey) => {
  return assetsMap[key];
};
