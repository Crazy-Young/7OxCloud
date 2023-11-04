<template>
    <div class="loading">
        <!-- 加载图片 -->
        <div class="loading-img"></div>
        <p class="loading-text">加载中...</p>
    </div>
</template>

<style lang="less" scoped>
.loading {
    margin: 20vh auto;
    width: 100%;
    box-sizing: content-box;
    text-align: center;

    .loading-img {
        // 使用雪碧图，步长为30px，逐帧播放
        background: url('../assets/loading.png') no-repeat;
        background-size: 30px auto;
        width: 30px;
        height: 30px;
        background-position: 0 0;
        animation: running 1s steps(60) infinite;
        margin: auto;

        // 动画
        @keyframes running {
            from {
                background-position: 0 0;
            }

            to {
                background-position: 0 -1800px;
            }
        }
    }

    .loading-text {
        color: @gray-3;
        margin-top: 10px;
    }
}
</style>