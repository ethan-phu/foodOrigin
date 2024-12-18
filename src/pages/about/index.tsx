import { View, Text, Image } from '@tarojs/components';
import { FC } from 'react';
import PageContainer from '@shared/components/PageContainer';
import { Cell } from '@nutui/nutui-react-taro';
import Taro from '@tarojs/taro';
import './index.scss';

const About: FC = () => {
  const navigateToPrivacy = () => {
    Taro.navigateTo({
      url: '/pages/privacy/index'
    });
  };

  return (
    <PageContainer title="关于我们" showBack>
      {/* Logo和应用名称 */}
      <View className="flex flex-col items-center pt-0 pb-6">
        <Image
          className="w-20 h-20 mb-3 rounded-2xl shadow-sm"
          src="https://example.com/logo.png"
          mode="aspectFit"
        />
        <Text className="text-lg font-medium">食知源</Text>
        <Text className="text-xs text-gray-500 mt-1">版本 1.0.0</Text>
      </View>

      {/* 关于信息列表 */}
      <View className="about-list mt-2">
        <Cell title="用户协议" onClick={() => {}} className="py-2.5" />
        <Cell title="隐私政策" onClick={navigateToPrivacy} className="py-2.5" />
        <Cell title="更新历史" onClick={() => {}} className="py-2.5" />
        <Cell title="开源许可" onClick={() => {}} className="py-2.5" />
      </View>

      {/* 底部版权信息 */}
      <View className="text-center text-gray-400 text-xs mt-6 mb-4">
        2024 食知源. All rights reserved.
      </View>
    </PageContainer>
  );
};

export default About;
