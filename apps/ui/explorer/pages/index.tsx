import React from 'react';
import { getAssetSrc, petApi } from '@course-explorer-monorepo/ui';
import { XStack, Input, Image } from 'tamagui';

export function Index() {
  React.useEffect(() => {
    petApi.getPetById({ petId: 1 }).then((pet) => console.log(pet));
  }, []);

  return (
    <XStack padding="$2">
      <Input placeholder="Search" />
      <Image
        height={100}
        width={100}
        src={getAssetSrc('icons/sample-icon')}
        alt="tes"
      />
    </XStack>
  );
}

export default Index;
