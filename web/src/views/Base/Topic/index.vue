<template>
    <!-- 滚动布局 -->
    <ScrollVideo :videoList="videoList" :isDataEnd="isDataEnd" @getVideoList="getVideoList"></ScrollVideo>
</template>

<script>
import { VideoListByTopic } from '@/api/video';
import videoList from '@/mixins/videoList';
import ScrollVideo from "@/views/ScrollVideo"
export default {
    name: "Topic",
    mixins: [videoList],
    components: {
        ScrollVideo
    },
    methods: {
        getVideoList() {
            return this.reqVideoList(() => VideoListByTopic(this.tid, this.latestTime))
        },
    },
    computed: {
        tid() {
            return this.$route.params.tid
        }
    }
}
</script>