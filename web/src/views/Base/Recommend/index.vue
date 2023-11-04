<template>
    <!-- 轮播布局 -->
    <SwiperVideo :getVideoFn="VideoListByRecommend"></SwiperVideo>
</template>

<script>
import SwiperVideo from '@/views/SwiperVideo';
import { VideoListByRecommend } from '@/api/video'
export default {
    name: "Friend",
    components: {
       SwiperVideo
    },
    methods: {
        VideoListByRecommend,
    },
};
</script>