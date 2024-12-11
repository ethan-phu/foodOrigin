import { View, Text } from "@tarojs/components";
import PageContainer from "@shared/components/PageContainer";
import { useEffect } from "react";

const Mine = () => {
  useEffect(() => {
    console.log("Mine Page Loaded");
  }, []);

  return (
    <PageContainer className="mine">
      <Text className="text-[#acc855] text-[32px] flex-wrap">我的</Text>
    </PageContainer>
  );
};

export default Mine;
