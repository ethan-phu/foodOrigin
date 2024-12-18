import { View, Text } from '@tarojs/components';
import { FC } from 'react';
import PageContainer from '@shared/components/PageContainer';
import './index.scss';

const PrivacyPolicy: FC = () => {
  return (
    <PageContainer title="隐私政策" showBack>
      <View className="privacy-policy p-4">
        <View className="section mb-6">
          <Text className="text-xl font-bold mb-4 block">隐私政策</Text>
          <Text className="text-gray-700 mb-2 block">
            本隐私政策介绍了食知源小程序（以下简称"我们"）如何收集、使用和保护您的个人信息。使用我们的服务即表示您同意本隐私政策的内容。
          </Text>
        </View>

        <View className="section mb-6">
          <Text className="text-lg font-bold mb-3 block">1. 信息收集</Text>
          <Text className="text-gray-700 mb-2 block">我们收集的信息包括：</Text>
          <View className="ml-4">
            <Text className="text-gray-700 mb-2 block">• 基本信息：微信授权的头像、昵称</Text>
            <Text className="text-gray-700 mb-2 block">• 设备信息：设备型号、操作系统版本</Text>
            <Text className="text-gray-700 mb-2 block">• 使用数据：搜索历史、浏览记录</Text>
          </View>
        </View>

        <View className="section mb-6">
          <Text className="text-lg font-bold mb-3 block">2. 信息使用</Text>
          <Text className="text-gray-700 mb-2 block">我们使用收集的信息：</Text>
          <View className="ml-4">
            <Text className="text-gray-700 mb-2 block">• 提供、维护和改进我们的服务</Text>
            <Text className="text-gray-700 mb-2 block">• 开发新的功能和服务</Text>
            <Text className="text-gray-700 mb-2 block">• 保护用户和公共安全</Text>
          </View>
        </View>

        <View className="section mb-6">
          <Text className="text-lg font-bold mb-3 block">3. 信息保护</Text>
          <Text className="text-gray-700 mb-2 block">
            我们采取适当的技术和组织措施来保护您的个人信息免受未经授权的访问、使用或披露。
          </Text>
        </View>

        <View className="section mb-6">
          <Text className="text-lg font-bold mb-3 block">4. 信息共享</Text>
          <Text className="text-gray-700 mb-2 block">
            除非经过您的同意，我们不会与第三方共享您的个人信息，但以下情况除外：
          </Text>
          <View className="ml-4">
            <Text className="text-gray-700 mb-2 block">• 遵守法律法规的要求</Text>
            <Text className="text-gray-700 mb-2 block">• 保护我们的合法权益</Text>
            <Text className="text-gray-700 mb-2 block">• 经您明确同意的其他情况</Text>
          </View>
        </View>

        <View className="section mb-6">
          <Text className="text-lg font-bold mb-3 block">5. 您的权利</Text>
          <Text className="text-gray-700 mb-2 block">您有权：</Text>
          <View className="ml-4">
            <Text className="text-gray-700 mb-2 block">• 访问您的个人信息</Text>
            <Text className="text-gray-700 mb-2 block">• 更正不准确的信息</Text>
            <Text className="text-gray-700 mb-2 block">• 删除您的账户</Text>
          </View>
        </View>

        <View className="section mb-6">
          <Text className="text-lg font-bold mb-3 block">6. 隐私政策更新</Text>
          <Text className="text-gray-700 mb-2 block">
            我们可能会不时更新本隐私政策。当我们进行重大更改时，我们会通过适当方式通知您。
          </Text>
        </View>

        <View className="section mb-6">
          <Text className="text-lg font-bold mb-3 block">7. 联系我们</Text>
          <Text className="text-gray-700 mb-2 block">
            如果您对本隐私政策有任何问题或建议，请通过小程序"关于我们"页面的反馈功能与我们联系。
          </Text>
        </View>
      </View>
    </PageContainer>
  );
};

export default PrivacyPolicy;
