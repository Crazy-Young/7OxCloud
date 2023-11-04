<template>
    <div id="video-cut">
        <div class="mask" @click="$emit('close')"></div>
        <div class="panel">
            <div class="panel-close" @click="$emit('close')"></div>
            <div class="panel__header">
                <div class="panel__header-title">截取视频封面</div>
            </div>

            <div class="panel-content">
                <div class="video">
                    <video :src="video" ref="video" crossOrigin="anonymous"></video>
                </div>
                <el-slider v-model="time" :min="min" :max="max" :step="0.1"></el-slider>
                <div>
                    <el-button type="primary" @click="submit">截取</el-button>
                    <el-button type="info" @click="$emit('close')">取消</el-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "VideoCut",
    props: {
        video: {
            type: String,
        }
    },
    data() {
        return {
            time: 0,
            max: 0,
            min: 0
        }
    },
    methods: {
        submit() {
            const canvas = document.createElement("canvas");
            const ctx = canvas.getContext("2d");
            const video = this.$refs.video;
            const pixelRatio = 0.5;
            canvas.width = video.videoWidth * pixelRatio;
            canvas.height = video.videoHeight * pixelRatio;
            ctx.drawImage(video, 0, 0, video.videoWidth, video.videoHeight, 0, 0, canvas.width, canvas.height);
            canvas.toBlob((blob) => {
                if (!blob) {
                    this.$message.error("截取视频封面失败");
                    this.$emit("close");
                    return;
                }
                blob.name = "截取封面.png";
                this.$emit("submit", blob);
                this.$emit("close");
            }, "image/png", 0.5);
        }
    },
    mounted() {
        const video = this.$refs.video;
        video.addEventListener("loadedmetadata", () => {
            this.max = video.duration;
        })
    },
    watch: {
        time() {
            this.$refs.video.currentTime = this.time;
        }
    }
};
</script>

<style lang="less" scoped>
#video-cut {
    position: fixed;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    z-index: 999;

    .mask {
        position: absolute;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        background-color: @mask;
    }

    .panel {
        position: absolute;
        top: 50%;
        left: 50%;
        display: flex;
        border-radius: 4px;
        min-height: 508px;
        padding: 0 40px;
        flex-direction: column;
        transform: translate(-50%, -50%);
        width: 560px;
        background: @gray-4;
        color: @white;
        padding-bottom: 40px;

        .panel-close {
            background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAJAAAACQCAYAAADnRuK4AAAACXBIWXMAACxLAAAsSwGlPZapAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAPTSURBVHgB7d29ThRRGIDh7wyLWGlFQoJ7A4rBBgs7CxPs8QqwsRMra3u8Am/AYGOnnR2yGxNZxM4KsgnQURrY43y7IQgL7PzPnHPeJzFCmGLIvMz/OSsCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoCxGPLGwsNSemjJrVuxzY+SOtWZPrN0UidZ7va19aQBdxygyqyaSZRHbtlaOxcqXJq1jWl4E9PDR4xfG2ncazuWfaUgnYld/b3d3pUaLi0sPrMinq9ZRDcS83vnZ2RDHOB/QcM/Tku83LaN/6SciK3VFNCkeNdobRc9c2xNF4rgokjeTltEN14o34P14Q0rFksSj9OfG2Im/S9M4H1C8D11OtFgNESWN54yev4ljnA8o6cY5W7aqiNLGo9Is2xTu74FEUp0zVBFRlnhG4itHx7gf0PAyOJ0yI8oez/AQtimOcT6g01vyYXgFk1IZEeWJZ7j3mYnWxTHOB/Sr292L70Ws1B1RrnhMvO7TdrW35d7NRG/uROc6dOS8T5Q7npasbHfrvdGZlTcBqToiCjke5VVAqsqIQo9HeReQqiIi4hnxMiBVZkTEc87bgFQZERHPRV4HpIqMiHjGeR+QKiKiaf2aeMYEEZDKG5H+TzzjgglI5XvUkIHn8aigAlKVRRRAPCq4gFTpEQUSjwoyIFVaRAHFo4INSBUeUWDxqKADUoVFFGA8KviAVO6IAo1HtQS52YHIyV8JEoewgg5hdQ9erAsn0QWeRIcYEZfxBV/GhxYRNxJLEFJEPMooSSgR8TC1RCFE5MPQ5kTyvgzWlMGLTRNEQEW8SdiEwYtNxCutN7l0h7nOwYtNxUv117nm8QQRXcSwnqtMeLZFROcYWHhZwgejRDTC0Ob/pXyqTkRMrnAu4ysZoUfE9C4q5/s8IUfkRUBNGDEaakReTDQeTQ033D1Jq+A3CXMPXmSi8erpRONNiEdtx3uQPHesJbIvxTE+TDT+RNIq8R3mXBHZZJOmN4kPE42n2/tU8AJ85oiMbYtjnA8o1UaqcPREloiy7LXq5sMhLNlE4zUMvUkdkUk/aXrdnA9oMC3vJ26gGsdtpYrIwYnGp8Rxh/3+8ezc/FF03af2NGDQ38FB/2hubv5bvC5P42/HL/HjdTy18nbnR8e5jzpwPiB1eNDfnW3Pb5iBfuaW6InoTPxv34r5GP9Vr/U6nT9SM40oXsevw3UUuSsakr7pOJDPcjt65WI8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQJX+ATc9oF5n57RZAAAAAElFTkSuQmCC);
            background-size: 100% 100%;
            cursor: pointer;
            height: 36px;
            position: absolute;
            right: 5px;
            top: 5px;
            width: 36px;
            z-index: 9999;
            filter: brightness(4);
        }

        .panel__header {
            background: transparent;
            border-radius: 4px 4px 0 0;
            width: 100%;
            margin-top: 36px;

            .panel__header-title {
                font-size: 18px;
                font-weight: 500;
                height: 34px;
                line-height: 34px;
            }

        }

        .panel-content {
            flex: 1;
            border-radius: 12px;
            min-height: 320px;
            display: flex;
            flex-direction: column;
            text-align: center;

            .video {
                height: 100%;
                width: 100%;
                max-height: 300px;
                margin: 40px 0 20px;

                video {
                    height: 100%;
                    max-height: 300px;
                    width: 100%;
                }
            }

            .el-slider {
                margin: 40px 40px 30px;

                /deep/ .el-slider__bar {
                    background-color: @red;
                }

                /deep/ .el-slider__button {
                    border-color: @red;
                }
            }

            .el-button {
                background-color: @red;
                border: unset;
                padding: 0 16px;
                height: 36px;
                line-height: 22px;
                min-width: 88px;
                font-size: 14px;

                &:disabled {
                    filter: brightness(0.3) !important;
                }

                &:hover {
                    filter: brightness(0.9);
                }

                &:active {
                    filter: brightness(0.8);
                }

                &--info {
                    background-color: @gray-3;
                }
            }
        }
    }


}
</style>