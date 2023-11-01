<template>
    <el-empty :description="msg"></el-empty>
</template>

<script>
import { Empty } from 'element-ui';
export default {
    name: 'Error',
    computed: {
        msg() {
            return `Can't find the page ${this.$route.fullPath} you are looking for!`
        }
    },
    components: {
        [Empty.name]: Empty
    }
}
</script>

<style lang="less" scoped>
.el-empty {
    position: absolute;
    inset: 0;
}
</style>