<template>
    <!-- 网格布局 -->
    <GridVideo :videoList="videoList" :isDataEnd="isDataEnd" @getVideoList="getVideoList"></GridVideo>
</template>

<script>
import { VideoListByCategory } from '@/api/video';
import GridVideo from "@/views/GridVideo";
import videoList from '@/mixins/videoList';
export default {
    name: "Channel",
    mixins: [videoList],
    components: {
        GridVideo
    },
    methods: {
        getVideoList() {
            this.reqVideoList(() => VideoListByCategory(this.cid, this.latestTime))
        }
    },
    computed: {
        cid() {
            return this.$route.params.cid
        }
    }
}
</script>