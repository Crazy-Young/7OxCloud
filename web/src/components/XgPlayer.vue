<template>
    <div :id="vid" class="xgplayer" @dblclick="navigateDetail" :style="`--bg: url(${coverUrl})`" @pointerdown.stop
        :class="{ up: ctrlpos !== 'bottom' }">
        <slot></slot>
        <slot name="title"></slot>
        <slot name="operation"></slot>
    </div>
</template>

<script>
import Xgplayer from "xgplayer";
export default {
    name: "XgPlayer",
    props: {
        vid: {
            type: String | Number | Object,
            default: "xgplayer",
        },
        playUrl: {
            type: String,
            default: () => "",
        },
        coverUrl: {
            type: String,
            default: () => ""
        },
        playsinline: {
            type: Boolean,
            default: true,
        },
        thumbnail: {
            type: Boolean,
            default: true,
        },
        muted: {
            type: Boolean,
            default: false,
        },
        width: {
            type: String,
            default: '100%'
        },
        height: {
            type: String,
            default: '100%'
        },
        ignores: {
            type: Array,
            default: () => ['cssFullscreen']
        },
        autoplay: {
            type: Boolean,
            default: false
        },
        autoHide: {
            type: Boolean,
            default: true
        },
        ctrlpos: {
            type: String,
            default: "normal",
        }
    },
    data() {
        return {
            player: null,
        };
    },
    methods: {
        initPlayer() {
            this.player = new Xgplayer({
                id: this.vid,
                url: this.playUrl,
                poster: this.coverUrl,
                playsinline: this.playsinline,
                width: this.width,
                height: this.height,
                volume: this.muted ? 0 : 1,
                loop: true,
                // thumbnail: this.thumbnail ? {
                //     pic_num: 44,
                //     width: 160,
                //     height: 90,
                //     col: 10,
                //     row: 10,
                //     urls: ["//lf9-cdn-tos.bytecdntp.com/cdn/expire-1-M/byted-player-videos/1.0.0/xgplayer-demo-thumbnail.jpg"],
                // } : false,
                ignores: this.ignores,
                autoplay: this.autoplay,
                controls: {
                    mode: this.ctrlpos,
                    autoHide: this.autoHide
                },
                closeVideoDblclick: true
            });
        },
        destroyPlayer() {
            this.player.destroy();
        },
        play() {
            this.player.play();
        },
        pause() {
            this.player.pause();
        },
        navigateDetail() {
            this.$router.push(`/video/${this.vid}`)
        }
    },
    mounted() {
        this.initPlayer();
    },
    beforeDestroy() {
        this.destroyPlayer();
    },
    watch: {
        src: {
            handler() {
                this.initPlayer();
            }
        }
    }
};
</script>

<style lang="less" scoped>
.xgplayer {
    z-index: 0;
    position: relative;

    &::after {
        background: var(--bg) 0% 0% / 100% no-repeat #000;
        background-position: center center;
        content: "";
        display: block;
        height: 100%;
        left: 0;
        position: absolute;
        top: 0;
        width: 100%;
        z-index: -1;
        background-size: cover;
        filter: blur(30px);
        transform: scale(1.1);
    }

    /deep/ .xgplayer-slider {
        display: none;
        border-radius: 6px;
        opacity: 1;
        background-color: @gray-4 ;
        box-sizing: content-box;
    }

    /deep/ .xg-options-list {
        opacity: 1;
        width: 70px;
        border-radius: 6px;
        background-color: @gray-4;

        .option-item {
            box-sizing: content-box;

            &:first-child {
                margin-top: 4px;
            }

            &:last-child {
                margin-bottom: 4px;
            }
        }
    }

    &.up {
        /deep/ .xgplayer-slider {
            bottom: 150%;
        }
    }

    /deep/ .xgplayer-playbackrate {
        padding-top: 20px;
        box-sizing: content-box;
        transform: translateY(-20px);
    }
}
</style>