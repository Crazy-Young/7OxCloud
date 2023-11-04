<template>
    <div class="home-container">
        <div class="home" id="home" ref="home">
            <template v-if="!loading || videoList.length>15">
                <VideoBox class="video" v-for="(item, index) in videoList" v-bind="item" :key="item.vid" ref="VideoBox"></VideoBox>
            </template>
        </div>
        <DataStatus :isDataEnd="isDataEnd" :length="videoList.length" ref="loading" @getVideoList="getVideoList"></DataStatus>
    </div>
</template>

<script>
import VideoBox from '@/components/VideoBox.vue';
import waterFall from '@/utils/waterfallLayout';
import { VideoListByTime } from '@/api/video';
import DataStatus from '@/components/DataStatus.vue';
import videoList from '@/mixins/videoList';
export default {
    name: "Home",
    mixins: [videoList],
    components: {
        VideoBox,
        DataStatus
    },
    data() {
        return {
            loading: true,
        }
    },
    methods: {
        // 设置瀑布流布局信息
        waterfallLayout() {
            waterFall(this.$refs.home, window.innerWidth < 800 ? 2 : window.innerWidth < 1280 ? 3 : window.innerWidth < 1560 ? 4 : 5, 20)
        },
        // 加载所有资源
        async loadAllSources() {
            this.loading = true
            for (let i = 0; i < this.videoList.length; i++) {
                if(this.videoList[i].isLoaded) continue
                const video = document.createElement('video');
                video.src = this.videoList[i].playUrl;
                await new Promise((resolve, reject) => {
                    // 监听视频加载完成
                    video.onloadedmetadata = () => {
                        // 获取视频尺寸
                        this.videoList[i].width = video.videoWidth;
                        this.videoList[i].height = video.videoHeight;
                        // 计算比例
                        this.videoList[i].ratio = this.videoList[i].height / this.videoList[i].width;
                        // 加载完成
                        this.videoList[i].isLoaded = true
                        resolve()
                    }
                    video.onerror = () => {
                        resolve()
                    }
                })
                // 移除video元素
                video.remove()
            }

            this.loading = false
            this.$nextTick(() => {
                // 重新布局
                this.waterfallLayout()
            })

        },
        async getVideoList() {
            // 获取视频列表
            await this.reqVideoList(() => VideoListByTime(this.latestTime))
            // 加载所有资源
            await this.loadAllSources()
        },
    },
    mounted() {
        // 监听窗口大小变化，重新布局
        window.addEventListener('resize', this.waterfallLayout)
    },
    beforeDestroy() {
        // 取消监听窗口大小变化
        window.removeEventListener('resize', this.waterfallLayout)
    }
}
</script>

<style lang="less" scoped>
.home-container {
    position: relative;

    .home {
        position: relative;
        margin: 0 30px;
        min-width: 480px;
        padding-bottom: 30px;
        min-height: 100%;

        .video {
            position: absolute;
        }
    }
}
</style>