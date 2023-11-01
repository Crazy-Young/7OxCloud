<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: "App",
  mounted() {
    this.$store.dispatch("UserInfo");
  }
};
</script>

<style lang="less">
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;

  a {
    text-decoration: none;
    color: @font;
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
}

#app {
  height: 100vh;
  overflow: hidden;
  color: @font;
  font-family: "PingFang SC", "Helvetica Neue", Helvetica, "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;

  .el-button {
    font-family: "PingFang SC", "Helvetica Neue", Helvetica, "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  }

}

.el-select-dropdown {
  background-color: @gray-4 !important;
  border: unset !important;

  .el-scrollbar__wrap {
    margin-right: -16px !important;
    margin-bottom: -16px !important;
    border-radius: inherit;
  }

  &__list {
    padding: 0 !important;
    border-radius: 4px;
    overflow: hidden;
    background-color: transparent;
    outline: unset;
    border: unset;
  }

  &__item {
    border: unset;
    color: @gray-1 !important;
    outline: unset;
    background-color: transparent;

    &.selected {
      color: @red !important;
      background-color: transparent !important;
    }

    &:hover,
    &.hover {
      background-color: @nav-bg-hover !important;
      filter: brightness(.9) !important;
    }
  }
}

.el-popper {
  .popper__arrow {
    display: none !important;
    border: none;
  }
}

// .xgplayer-enter {
//   display: none !important;
// }

.xgplayer-slider {
  bottom: 150% !important;
  border-radius: 6px !important;
  opacity: 1 !important;
  background-color: @gray-4 !important;
  box-sizing: content-box;
}

.xg-options-list {
  bottom: 150% !important;
  opacity: 1 !important;
  width: 70px !important;
  border-radius: 6px !important;
  background-color: @gray-4 !important;

  .option-item {
    box-sizing: content-box;

    &:first-child {
      margin-top: 4px !important;
    }

    &:last-child {
      margin-bottom: 4px !important;
    }
  }
}
</style>
