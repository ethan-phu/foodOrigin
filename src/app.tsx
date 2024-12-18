import { useEffect } from "react";
import { Provider } from "mobx-react";
import Taro from "@tarojs/taro";
import store, { userStore } from "@shared/store";
import "@nutui/nutui-react-taro/dist/style.css";
import "./app.scss";

const App = (props) => {
  useEffect(() => {
    if (process.env.TARO_ENV === "weapp") {
      // 检查登录状态并自动登录
      const checkAndLogin = async () => {
        try {
          await Taro.checkSession();
          console.log('Session is valid');
          // 如果会话有效且有 token，尝试刷新 token
          if (userStore.token) {
            try {
              await userStore.refreshToken();
              console.log('Token refreshed successfully');
              // 获取用户信息
              await userStore.fetchUserInfo();
              console.log('User info fetched successfully');
            } catch (error) {
              console.error('Failed to refresh token or fetch user info:', error);
              // token 刷新失败，可能是过期了，清除登录状态
              userStore.logout();
            }
          }
        } catch (error) {
          console.log('Session expired or not exist');
          // 会话过期，清除登录状态
          userStore.logout();
        }
      };

      checkAndLogin();

      // 检测新版本
      if (Taro.getUpdateManager) {
        const updateManager = Taro.getUpdateManager();
        updateManager.onCheckForUpdate((res) => {
          // 请求完新版本信息的回调
          res.hasUpdate && console.warn("新版本提示");
        });

        updateManager.onUpdateReady(() => {
          Taro.showModal({
            title: "更新提示",
            content: "新版本已经准备好，是否重启应用？",
            success: (res) => {
              if (res.confirm) {
                // 新的版本已经下载好，调用 applyUpdate 应用新版本并重启
                updateManager.applyUpdate();
              }
            },
          });
        });

        updateManager.onUpdateFailed(() => {
          // 新版本下载失败
          Taro.showModal({
            title: "更新提示",
            content: "新版本下载失败，请检查网络后重试",
            showCancel: false,
          });
        });
      }
    }
  }, []);

  return <Provider store={store}>{props.children}</Provider>;
};

export default App;
