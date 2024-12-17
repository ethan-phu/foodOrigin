import { useEffect } from "react";
import { Provider } from "mobx-react";
import Taro from "@tarojs/taro";
import store, { userStore } from "@shared/store";
import { ApiConfig } from "@/config";
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
          // 如果会话有效，尝试获取用户信息
          if (userStore.token) {
            try {
              await userStore.fetchUserInfo();
              console.log('User info fetched successfully');
            } catch (error) {
              console.error('Failed to fetch user info:', error);
            }
          }
        } catch (error) {
          console.log('Session expired or not exist, trying to login');
          try {
            const { code } = await Taro.login();
            console.log('Login successful, code:', code);
            
            const res = await Taro.request({
              url: `${ApiConfig.base_url}/wechat/login`,
              method: "POST",
              data: {
                code,
              },
            });
            
            if (res.statusCode === 200 && res.data.err_code === 0) {
              userStore.setToken(res.data.data.token);
              // 登录成功后获取用户信息
              try {
                await userStore.fetchUserInfo();
                console.log('User info fetched successfully');
              } catch (error) {
                console.error('Failed to fetch user info:', error);
              }
            } else {
              console.error("Login failed:", res.data.message);
            }
          } catch (loginError) {
            console.error('Login failed:', loginError);
          }
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
            success(res) {
              if (res.confirm) {
                updateManager.applyUpdate(); // 新的版本已经下载好，调用 applyUpdate 应用新版本并重启
              }
            },
          });
        });
        updateManager.onUpdateFailed(() => {
          // 新的版本下载失败
          Taro.showModal({
            title: "已经有新版本了哟~",
            content: "新版本已经上线啦~，请您删除当前小程序，重新搜索打开哟~",
          });
        });
      } else {
        Taro.showModal({
          title: "提示",
          content:
            "当前微信版本过低，无法使用该功能，请升级到最新微信版本后重试。",
        });
      }
    }
  }, []);

  return <Provider store={store}>{props.children}</Provider>;
};

export default App;
