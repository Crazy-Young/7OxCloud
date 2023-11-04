<template>
    <!-- 错误路由 -->
    <el-empty :description="msg">
        <el-button @click="$router.replace('/')">返回首页</el-button>
    </el-empty>
</template>

<script>
export default {
    name: 'Error',
    computed: {
        msg() {
            return `对不起! 找不到你要跳转的 ${this.$route.fullPath} 页面!`
        }
    },
}
</script>

<style lang="less" scoped>
.el-empty {
    position: absolute;
    inset: 0;

    .el-button {
        background-color: @red;
        border: unset;
        color: @white;
    }
}
</style>