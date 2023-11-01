<template>
    <div class="home-container">
        <div class="home" id="home" ref="home">
            <template v-if="!loading">
                <VideoBox class="video" v-for="(item, index) in videoList" v-bind="item" :key="item.vid" ref="VideoBox"
                    @mouseenter.native="handleMouseEnter(index)" @mouseleave.native="handleMouseLeave(index)"></VideoBox>
            </template>
        </div>
        <el-collapse-transition>
            <div class="loading" v-if="!isDataEnd" ref="loading">
                <p class="loading-text"><i class="el-icon-loading icon"></i>加载中...</p>
            </div>
        </el-collapse-transition>
    </div>
</template>

<script>
import VideoBox from '@/components/VideoBox.vue';
import waterFall from '@/utils/waterfallLayout';
import { VideoListByTime } from '@/api/video';
export default {
    name: "Home",
    components: {
        VideoBox
    },
    data() {
        return {
            loading: true,
            timer: null,
            videoList: [],
            latestTime: null,
            isDataEnd: false,
            ob: null
        }
    },
    methods: {
        waterfallLayout() {
            waterFall(this.$refs.home, window.innerWidth < 800 ? 2 : window.innerWidth < 1280 ? 3 : window.innerWidth < 1560 ? 4 : 5, 30)
        },
        async loadAllSources() {
            for (let i = 0; i < this.videoList.length; i++) {
                const video = document.createElement('video');
                video.src = this.videoList[i].playUrl;
                await new Promise((resolve, reject) => {
                    video.onloadedmetadata = () => {
                        this.videoList[i].width = video.videoWidth;
                        this.videoList[i].height = video.videoHeight;
                        this.videoList[i].ratio = this.videoList[i].height / this.videoList[i].width;
                        resolve()
                    }
                    video.onerror = () => {
                        resolve()
                    }
                })
                video.remove()
            }

            this.loading = false
            this.$nextTick(() => {
                this.waterfallLayout()
            })

        },
        handleMouseEnter(index) {
            this.timer = setTimeout(() => {
                this.$refs.VideoBox[index].play()
            }, 100)
        },
        handleMouseLeave(index) {
            clearTimeout(this.timer)
            this.$refs.VideoBox[index].pause()
        },
        async getVideoList() {
            try {
                const res = await VideoListByTime(this.latestTime)
                if (res.status !== 200) {
                    this.$message.error(res.data.msg)
                    return
                }
                if (res.data.data.videos.length === 0) {
                    this.isDataEnd = true
                    return
                }
                this.videoList.push(...res.data.data.videos)
                this.latestTime = res.data.data.nextTime
                await this.loadAllSources()
                this.loading = false
            } catch (error) {
                this.$message.error(error)
                this.loading = false
            }
        }
    },
    mounted() {
        window.addEventListener('resize', this.waterfallLayout)

        this.ob = new IntersectionObserver((entries) => {
            if (entries[0].isIntersecting) {
                this.getVideoList()
            }
        }, {
            threshold: 0.5,
        })

        this.ob.observe(this.$refs.loading)
    },
    beforeDestroy() {
        window.removeEventListener('resize', this.waterfallLayout)
        this.ob.unobserve(this.$refs.loading)
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

    .loading {
        position: absolute;
        bottom: 0;
        transform: translateY(100%);
        padding: 100px 0;
        margin: auto;
        width: 100%;
        box-sizing: content-box;
        text-align: center;

        .icon {
            margin-right: 5px;
        }
    }
}
</style>