import { View, Text, Button } from "@tarojs/components";
import { observer } from "mobx-react";
import PageContainer from "@shared/components/PageContainer";
import { useState } from "react";
import { userStore } from "@shared/store";
import { Cell, Avatar } from "@nutui/nutui-react-taro";
import Taro from "@tarojs/taro";
import "./index.scss";

const Mine = observer(() => {
  const [loading, setLoading] = useState(false);

  const handleLogin = async () => {
    try {
      setLoading(true);
      const { userInfo } = await Taro.getUserProfile({
        desc: '用于完善会员资料',
        lang: 'zh_CN'
      });
      try {
        await userStore.login(userInfo.avatarUrl, userInfo.nickName);
      } catch (error) {
        console.error('Login failed:', error);
        Taro.showToast({ title: '登录失败', icon: 'error' });
        return;
      }
      userStore.updateUserInfo({
        nickName: userInfo.nickName,
        avatarUrl: userInfo.avatarUrl
      });
      Taro.showToast({ title: '登录成功', icon: 'success' });
    } catch (error) {
      console.error('Login failed:', error);
      Taro.showToast({ title: '登录失败', icon: 'error' });
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    userStore.logout();
    Taro.showToast({
      title: '退出登录成功',
      icon: 'success',
      duration: 2000
    });
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
        <Cell title="浏览历史" />
        <Cell title="意见反馈" />
        <Cell title="关于我们" onClick={navigateToAbout} />
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
