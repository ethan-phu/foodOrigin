export default defineAppConfig({
  pages: [
    'pages/home/index',
    'pages/category/index',
    'pages/mine/index'
  ],
  window: {
    backgroundTextStyle: "light",
    navigationBarBackgroundColor: "#fff",
    navigationBarTitleText: "食知源",
    navigationBarTextStyle: "black",
  },
  tabBar: {
    color: "#666666",
    selectedColor: "#acc855",
    backgroundColor: "#ffffff",
    list: [
      {
        pagePath: "pages/home/index",
        text: "首页",
        iconPath: "shared/assets/images/tabbar/home.png",
        selectedIconPath: "shared/assets/images/tabbar/home-active.png"
      },
      {
        pagePath: "pages/category/index",
        text: "分类",
        iconPath: "shared/assets/images/tabbar/category.png",
        selectedIconPath: "shared/assets/images/tabbar/category-active.png"
      },
      {
        pagePath: "pages/mine/index",
        text: "我的",
        iconPath: "shared/assets/images/tabbar/mine.png",
        selectedIconPath: "shared/assets/images/tabbar/mine-active.png"
      }
    ]
  },
  lazyCodeLoading: "requiredComponents",
  style: "v2",
  permission: {
    'scope.userLocation': {
      desc: '你的位置信息将用于获取当地天气信息' 
    }
  },
  requiredPrivateInfos: [
    'getLocation'
  ]
})
