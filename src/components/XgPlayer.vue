<template>
    <div :id="vid" class="xgplayer" :style="`--bg: url(${coverUrl})`" @pointerdown.stop>
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
            default: () => []
        },
        autoplay: {
            type: Boolean,
            default: false
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
                thumbnail: this.thumbnail ? {
                    pic_num: 44,
                    width: 160,
                    height: 90,
                    col: 10,
                    row: 10,
                    urls: ["//lf9-cdn-tos.bytecdntp.com/cdn/expire-1-M/byted-player-videos/1.0.0/xgplayer-demo-thumbnail.jpg"],
                } : false,
                danmu: {
                    comments: [
                        {
                            duration: 15000,
                            id: "1",
                            start: 3000,
                            txt: "长弹幕长弹幕长弹幕长弹幕长弹幕",
                            style: {
                                //弹幕自定义样式
                                color: "#ff9500",
                                fontSize: "20px",
                                border: "solid 1px #ff9500",
                                borderRadius: "50px",
                                padding: "5px 11px",
                                backgroundColor: "rgba(255, 255, 255, 0.1)",
                            },
                        },
                    ],
                    area: {
                        start: 0,
                        end: 1,
                    },
                },
                whitelist: [""],
                ignores: this.ignores,
                autoplay: this.autoplay,
                controls: {
                    mode: this.ctrlpos,
                }
            });

            // this.player.emit('resourceReady', [{ name: '超清', url: '//sf1-cdn-tos.huoshanstatic.com/obj/media-fe/xgplayer_doc_video/mp4/xgplayer-demo-720p.mp4' }, { name: '高清', url: '//sf1-cdn-tos.huoshanstatic.com/obj/media-fe/xgplayer_doc_video/mp4/xgplayer-demo-480p.mp4' }, { name: '标清', url: '//sf1-cdn-tos.huoshanstatic.com/obj/media-fe/xgplayer_doc_video/mp4/xgplayer-demo-360p.mp4' }]);
        },
        destroyPlayer() {
            this.player.destroy();
        },
        play() {
            this.player.play();
        },
        pause() {
            this.player.pause();
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
        content: "";
        display: block;
        height: 100%;
        left: 0;
        position: absolute;
        top: 0;
        width: 100%;
        z-index: -1;
        background-size: cover;
        filter: blur(10px);
        transform: scale(1.1);
    }
}
</style>