<template>
    <div class="topic-container">
        <div class="topic" id="topic" ref="topic">
            <template v-if="!loading">
                <VideoBox class="video" v-for="(item, index) in videoList" v-bind="item" :key="item.vid" ref="VideoBox"
                    :waterfall="false" @mouseenter.native="handleMouseEnter(index)"
                    @mouseleave.native="handleMouseLeave(index)"></VideoBox>
            </template>
        </div>
        <el-collapse-transition>
            <div class="loading" ref="loading">
                <p class="loading-text"><i class="el-icon-loading icon"></i>加载中...</p>
            </div>
        </el-collapse-transition>
    </div>
</template>

<script>
import VideoBox from '@/components/VideoBox.vue';
import waterFall from '@/utils/waterfallLayout';
import { VideoListByTopic } from '@/api/video';
export default {
    name: "Topic",
    components: {
        VideoBox
    },
    data() {
        return {
            loading: true,
            timer: null,
            videoList: [],
            latestTime: null,
            ob: null
        }
    },
    methods: {
        handleMouseEnter(index) {
            this.timer = setTimeout(() => {
                this.$refs.VideoBox[index].play()
            }, 300)
        },
        handleMouseLeave(index) {
            clearTimeout(this.timer)
            this.$refs.VideoBox[index].pause()
        },
        async getVideoList() {
            const res = await VideoListByTopic(this.tid, this.latestTime)
            if (res.status !== 200) {
                this.$message.error(res.data.msg)
                return
            }
            this.videoList.push(...res.data.data.videos)
            this.latestTime = res.data.data.nextTime
            this.loading = false
        }
    },
    mounted() {
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
        this.ob.unobserve(this.$refs.loading)
    },
    computed: {
        tid() {
            return this.$route.params.tid
        }
    }
}
</script>

<style lang="less" scoped>
.topic-container {
    position: relative;

    .topic {
        position: relative;
        margin: 0 30px;
        min-width: 480px;
        padding-bottom: 30px;
        min-height: 100%;
        display: grid;
        gap: 20px;
        grid-template-columns: repeat(5, 1fr);
        grid-template-rows: auto;

        @media screen and (max-width: 1200px) {
            grid-template-columns: repeat(3, 1fr);
        }

        @media screen and (max-width: 900px) {
            grid-template-columns: repeat(2, 1fr);
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