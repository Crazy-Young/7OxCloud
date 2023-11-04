<template>
    <div class="data-status">
        <!-- 加载中：数据没有加载完毕 -->
        <Loading v-if="!isDataEnd || loading" ref="loading"></Loading>
        <!-- 加载完毕：数据长度为0 -->
        <div class="no-data" v-else-if="!length">
            <el-empty description="暂无相关视频"></el-empty>
        </div>
        <!-- 所有数据加载完毕：数据不为零 -->
        <p class="no-more" v-else>暂时没有更多了</p>
    </div>
</template>

<script>
import Loading from './Loading.vue'
export default {
    name: "DataStatus",
    components: {
        Loading
    },
    data() {
        return {
            observe: null
        }
    },
    props: {
        // 数据长度
        length: {
            type: Number,
            default: 0
        },
        // 所有内容是否加载完毕
        isDataEnd: {
            type: Boolean,
            default: false
        },
        loading: {
            type: Boolean,
            default: false
        }
    },
    mounted() {
        // 定义监听行为
        this.observe = new IntersectionObserver((entries) => {
            // 只有一个元素被监听，所以entries[0]触发isIntersecting就更新数据
            if (entries[0].isIntersecting) {
                this.$emit("getVideoList")
            }
        }, {
            // 观察的距交叉的距离
            threshold: 0.5,
        })

        // 监听Loading组件是否出现在视口(包含第一次加载的情况)
        this.observe.observe(this.$refs.loading.$el)
    },
    beforeDestroy() {
        // 取消监听
        this.$refs.loading && this.observe.unobserve(this.$refs.loading.$el)
    }

}
</script>

<style lang="less" scoped>
.data-status {
    position: relative;

    .loading {
        margin: 20vh 0;
    }

    .no-data {
        margin: 20vh auto;
    }

    .no-more {
        margin: 20vh auto;
        text-align: center;
        color: @gray-3;
    }
}
</style>