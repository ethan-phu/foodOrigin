import { Text } from "@tarojs/components";
import { useEffect } from "react";
import PageContainer from "@shared/components/PageContainer";
import observer, { GlobalStore } from "@shared/store";
import "./index.scss";

const Index = () => {
  useEffect(() => {
    console.log(GlobalStore.customerName);
  }, []);
  return (
    <PageContainer className="index">
      <Text className="text-[#acc855] text-[32px] flex-wrap">你好这个是我的第一个 react app</Text>
    </PageContainer>
  );
};

export default observer(Index);
