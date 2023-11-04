<template>
    <!-- 轮播布局 -->
    <SwiperVideo :getVideoFn="VideoListByFriend"></SwiperVideo>
</template>

<script>
import SwiperVideo from '@/views/SwiperVideo';
import { VideoListByFriend } from '@/api/video'
export default {
    name: "Friend",
    components: {
       SwiperVideo
    },
    methods: {
        VideoListByFriend,
    },
};
</script>