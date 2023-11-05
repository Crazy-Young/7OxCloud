<template>
    <GridVideo :videoList="videoList" :isDataEnd="isDataEnd" @getVideoList="getVideoList"></GridVideo>
</template>

<script>
import { VideoListByHot } from '@/api/video';
import GridVideo from "@/views/GridVideo";
import videoList from '@/mixins/videoList';

export default {
    name: "Hot",
    mixins: [videoList],
    components: {
        GridVideo
    },
    methods: {
        getVideoList() {
            this.reqVideoList(() => VideoListByHot())
        }
    },
}
</script>