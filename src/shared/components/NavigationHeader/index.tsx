import Taro, { getSystemInfoSync, getMenuButtonBoundingClientRect } from "@tarojs/taro";
import { Text, View } from "@tarojs/components";
import router from "@shared/router";
import IconFont from "../IconFont";
import "./index.scss";

interface IPros {
  title?: string;
  customBacKfn?: () => void;
  className?: string;
}

interface NavigationBarInfo {
  navigationBarHeight: number;
  navigationContentHeight: number;
  menuButtonHeight: number;
  navigationPaddding: number;
  statusBarHeight: number;
  menuButtonWidth: number;
}

// 通过获取系统信息计算导航栏高度
const getNavigationBarInfo = (): NavigationBarInfo => {
  // 系统信息
  const systemInfo = getSystemInfoSync();
  // 胶囊按钮位置信息
  const menuButtonInfo = getMenuButtonBoundingClientRect();
  let navigationContentHeight = 40;
  // 胶囊导航栏高度
  navigationContentHeight =
    (menuButtonInfo.top - systemInfo.statusBarHeight) * 2 +
    menuButtonInfo.height;
  // 顶部手机状态栏高度
  const { statusBarHeight } = systemInfo;

  return {
    navigationBarHeight: statusBarHeight + navigationContentHeight,
    navigationContentHeight,
    menuButtonHeight: menuButtonInfo.height,
    navigationPaddding: systemInfo.windowWidth - menuButtonInfo.right,
    statusBarHeight: systemInfo.statusBarHeight??0,
    menuButtonWidth: menuButtonInfo.width,
  };
};

const NavigationHeader: React.FC<IPros> = (props) => {
  const { title, customBacKfn, className = '' } = props;

  const {
    statusBarHeight,
    navigationBarHeight,
    navigationContentHeight,
    menuButtonHeight,
    navigationPaddding,
    menuButtonWidth,
  } = getNavigationBarInfo();

  const onBackClick = () => {
    if (customBacKfn) {
      customBacKfn();
    } else if (Taro.getCurrentPages().length > 1) {
      router.navigateBack();
    } else {
      router.reLaunch("/pages/newHome");
    }
  };

  const onBackHome = () => {
    router.reLaunch("/pages/newHome");
  };

  const navStyle = {
    height: `${navigationBarHeight}px`,
    padding: `0 ${navigationPaddding}px`
  };

  const navbarStyle = {
    height: `${navigationContentHeight}px`,
    top: `${statusBarHeight}px`
  };

  const backIconStyle = {
    width: `${menuButtonWidth}px`,
    height: `${menuButtonHeight}px`,
    borderRadius: `${menuButtonHeight / 2}px`
  };

  const rightIconStyle = {
    width: `${menuButtonWidth}px`,
    height: `${menuButtonHeight}px`
  };

  return (
    <View className={`nav_home_bar ${className}`} style={navStyle}>
      <View className="navbar" style={navbarStyle}>
        <View className="back_icon" style={backIconStyle}>
          <View className="icon_item" onClick={onBackClick}>
            <IconFont icon="icon-fanhui" />
          </View>
          <View className="icon_item" onClick={onBackHome}>
            <IconFont icon="icon-shouye_xuanzhong" />
          </View>
        </View>
        {title && <Text className="nav_title">{title}</Text>}
        <View className="right_icon" style={rightIconStyle}></View>
      </View>
    </View>
  );
};

export default NavigationHeader;
