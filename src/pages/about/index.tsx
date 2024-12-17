import { View, Text, Image } from "@tarojs/components";
import { observer } from "mobx-react";
import PageContainer from "@shared/components/PageContainer";
import { Cell } from "@nutui/nutui-react-taro";
import "./index.scss";

const About = observer(() => {
  return (
    <PageContainer className="about bg-gray-100 min-h-screen">
      {/* Logo 和应用信息 */}
      <View className="app-info bg-white p-6 mb-4 text-center">
        <Image 
          className="app-logo mx-auto mb-4"
          src="/assets/images/logo.png"
          mode="aspectFit"
        />
        <Text className="app-name block text-xl font-bold mb-2">
          食知源
        </Text>
        <Text className="app-version text-gray-500">
          版本 1.0.0
        </Text>
      </View>

      {/* 关于信息列表 */}
      <View className="about-list">
        <Cell title="用户协议" isLink onClick={() => {}} />
        <Cell title="隐私政策" isLink onClick={() => {}} />
        <Cell title="更新历史" isLink onClick={() => {}} />
        <Cell title="开源许可" isLink onClick={() => {}} />
      </View>

      {/* 底部版权信息 */}
      <View className="copyright text-center text-gray-500 py-8">
        <Text className="block">© 2024 食知源</Text>
        <Text className="block mt-1">All Rights Reserved</Text>
      </View>
    </PageContainer>
  );
});

export default About;
