<template>
  <div id="app">
    <router-view></router-view>
    <Login v-if="isLogin" @close="toggleLogin"></Login>
  </div>
</template>

<script>
import Login from "@/components/Login";
export default {
  name: "App",
  components: {
    Login
  },
  data() {
    return {
      isLogin: false
    };
  },
  methods: {
    // 切换登录
    toggleLogin() {
      this.isLogin = !this.isLogin;
    }
  },
  mounted() {
    // 获取用户信息
    this.$store.dispatch("UserInfo");
    // 获取分类
    this.$store.dispatch("getCategoryList")
    // 总线
    this.$bus.$on("toggleLogin", this.toggleLogin);
  }
};
</script>

<style lang="less">
@import url("~@/less/element.less");
@import url("@/less/index.less");

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  word-break: break-word;

  a {
    text-decoration: none;
    color: unset;
  }

  li {
    list-style: none;
  }
}

@font-face {
  font-family: 'PingFang SC';
  src: url('./assets/font/PingFang\ SC.ttf');
}

body {
  background-color: @bg;
  background-image:
    radial-gradient(at -2% -2%, #21657eb0 0px, transparent 10%),
    radial-gradient(at 6% 5%, #6f285c9a 0px, transparent 10%);
    background-repeat: no-repeat;
    
}

#app {
  height: 100vh;
  color: @font;
  font-family: "PingFang SC", "Helvetica Neue", Helvetica, "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
}
</style>
