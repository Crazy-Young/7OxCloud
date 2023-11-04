<template>
    <!-- 视频布局: 表格 -->
    <div class="grid-video">
        <div class="channel" id="channel" ref="channel">
            <VideoBox class="video" v-for="(item, index) in videoList" v-bind="item" :key="item.vid" ref="VideoBox"
                :waterfall="false"></VideoBox>
        </div>
        <DataStatus :isDataEnd="isDataEnd" :length="videoList.length" ref="loading" @getVideoList="getVideoList"></DataStatus>
    </div>
</template>

<script>
import VideoBox from '@/components/VideoBox.vue';
import DataStatus from '@/components/DataStatus.vue';
export default {
    name: "GridVideo",
    props: {
        videoList: {
            type: Array,
            default: () => []
        },
        isDataEnd: {
            type: Boolean,
            default: false
        }
    },
    components: {
        VideoBox,
        DataStatus
    },
    methods: {
        getVideoList() {
            this.$emit("getVideoList");
        }
    },
}
</script>

<style lang="less" scoped>
.grid-video {
    position: relative;

    .channel {
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

}
</style>