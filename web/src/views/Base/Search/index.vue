<template>
    <!-- 滚动布局 -->
    <ScrollVideo :videoList="videoList" :isDataEnd="isDataEnd" @getVideoList="getVideoList"></ScrollVideo>
</template>

<script>
import { VideoListBySearch } from '@/api/video';
import videoList from '@/mixins/videoList';
import ScrollVideo from "@/views/ScrollVideo"
export default {
    name: "Search",
    mixins: [videoList],
    components: {
        ScrollVideo
    },
    methods: {
        getVideoList() {
            if (!this.keyword.trim()) {
                this.isDataEnd = true
                return
            }
            this.reqVideoList(() => VideoListBySearch(this.keyword, this.latestTime))
        },
    },
    computed: {
        keyword() {
            return this.$route.query.keyword
        }
    }
}
</script>