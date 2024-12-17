import { View, Text, Button } from "@tarojs/components";
import { observer } from "mobx-react";
import PageContainer from "@shared/components/PageContainer";
import { useState } from "react";
import { userStore } from "@shared/store";
import { Cell, Avatar, Toast } from "@nutui/nutui-react-taro";
import Taro from "@tarojs/taro";
import "./index.scss";

const Mine = observer(() => {
  const [loading, setLoading] = useState(false);

  const handleLogin = async () => {
    try {
      setLoading(true);
      await userStore.login();
      Toast.show('登录成功');
    } catch (error) {
      console.error('Login failed:', error);
      Toast.show('登录失败');
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    userStore.logout();
    Toast.show('已退出登录');
  };

  const navigateToAbout = () => {
    Taro.navigateTo({
      url: '/pages/about/index'
    });
  };

  return (
    <PageContainer className="mine bg-gray-100 min-h-screen">
      <View className="user-info bg-white p-4 mb-4">
        {userStore.isLoggedIn ? (
          <View className="flex items-center">
            <Avatar 
              size="large"
              src={userStore.userInfo?.avatar || ''} 
              className="mr-4"
            />
            <View className="flex-1">
              <Text className="text-lg font-bold">
                {userStore.userInfo?.nickname || userStore.userInfo?.name || '未设置昵称'}
              </Text>
              <Text className="text-gray-500 text-sm block mt-1">
                ID: {userStore.userInfo?.id}
              </Text>
            </View>
          </View>
        ) : (
          <View className="flex items-center justify-between py-2">
            <Avatar size="large" className="mr-4" />
            <View className="flex-1">
              <Button
                className="login-btn"
                loading={loading}
                onClick={handleLogin}
              >
                点击登录
              </Button>
            </View>
          </View>
        )}
      </View>

      <View className="menu-list">
        <Cell title="浏览历史" isLink />
        <Cell title="意见反馈" isLink />
        <Cell title="关于我们" isLink onClick={navigateToAbout} />
        {userStore.isLoggedIn && (
          <Cell
            title="退出登录"
            onClick={handleLogout}
            className="text-red-500"
          />
        )}
      </View>
    </PageContainer>
  );
});

export default Mine;
