import { View, Text } from "@tarojs/components";
import PageContainer from "@shared/components/PageContainer";
import { useEffect } from "react";

const Category = () => {
  useEffect(() => {
    console.log("Category Page Loaded");
  }, []);

  return (
    <PageContainer className="category">
      <Text className="text-[#acc855] text-[32px] flex-wrap">分类</Text>
    </PageContainer>
  );
};

export default Category;
