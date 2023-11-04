<template>
    <!-- 轮播布局 -->
    <SwiperVideo :getVideoFn="VideoListByFollow"></SwiperVideo>
</template>

<script>
import SwiperVideo from '@/views/SwiperVideo';
import { VideoListByFollow } from '@/api/video'
export default {
    name: "Follow",
    components: {
       SwiperVideo
    },
    methods: {
        VideoListByFollow,
    },
};
</script>