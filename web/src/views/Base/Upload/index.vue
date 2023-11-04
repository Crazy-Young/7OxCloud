<template>
    <div class="container">
        <h1 class="title">发布视频</h1>
        <el-form class="upload-container" ref="form" @submit.prevent="submit" :model="form" :rules="rules">
            <el-form-item prop="video">
                <el-upload class="upload upload-video" ref="video" :file-list="form.videos" :disabled="isSubmit"
                    :http-request="handleVideoUpload" :on-change="handleVideoChange" :on-remove="handleVideoRemove" drag
                    action="#" accept="video/*" :auto-upload="true" :limit="1">
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">将视频拖到此处，或<em>点击上传</em></div>
                </el-upload>
            </el-form-item>

            <el-form-item prop="desc">
                <template #label>
                    <span>作品描述</span>
                    <i class="count">{{ form.desc.length }}/300</i>
                </template>
                <el-input :disabled="isSubmit" type="textarea" rows="5" autosize v-model="form.desc" resize="none"
                    placeholder="写一个视频描述，让更多人看到" maxlength="300">
                </el-input>
            </el-form-item>

            <el-form-item prop="topics">
                <template #label>
                    <span>话题</span>
                </template>
                <el-select :disabled="isSubmit" width="200" v-model="form.topics" placeholder="自定义话题或选择" multiple
                    :multiple-limit="3" allow-create filterable @keyup.enter.native="addTopic">
                    <el-option v-for="item in topics" :key="item.id" :label="item.description" :value="item.description">
                    </el-option>
                </el-select>
            </el-form-item>

            <el-form-item prop="category">
                <template #label>
                    <span>类别</span>
                </template>
                <el-select :disabled="isSubmit" v-model="form.category" placeholder="选择类别">
                    <el-option v-for="item in categoryList" :key="item.id" :label="item.name" :value="item.id">
                    </el-option>
                </el-select>
            </el-form-item>

            <el-form-item prop="imgs">
                <template #label>
                    <span>设置封面</span>
                    <el-button class="cut" @click="isCut = true"
                        :disabled="!form.videos[0]?.url || isSubmit">视频截帧</el-button>
                </template>
                <el-upload :disabled="isSubmit" class="upload upload-pic" ref="img" :on-remove="handleImgRemove"
                    :on-change="handleImgChange" :http-request="handleImgUpload" :file-list="form.imgs" drag action="#"
                    accept="image/*" :auto-upload="true" :limit="1">
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">将封面拖到此处，或<em>点击上传</em></div>
                </el-upload>
            </el-form-item>

            <el-button class="submit" type="submit" @click="submit"
                :disabled="!form.videos.length || !form.imgs.length || isSubmit" :loading="isSubmit">发布</el-button>
        </el-form>

        <VideoCut v-if="isCut" @close="isCut = false" @submit="getImg" :video="form.videos[0]?.url"></VideoCut>
    </div>
</template>

<script>
import VideoCut from "./VideoCut";
import { DeleteFile, UploadFile } from "@/api/base";
import { VideoPublish, VideoTopicList } from "@/api/video";
import { mapState } from "vuex";
import getSensitive from "@/utils/getSensitive";
export default {
    name: "Upload",
    data() {
        return {
            isCut: false,
            isSubmit: false,
            form: {
                desc: "",
                videos: [],
                imgs: [],
                topics: [],
                category: ''
            },
            rules: {
                desc: [
                    { required: true, message: "请输入作品描述", trigger: "blur" },
                    { min: 0, max: 300, message: "长度在 0 到 300 个字符", trigger: "blur" }
                ],
                category: [
                    { required: true, message: "请选择类别", trigger: "change" }
                ],
                videos: [
                    { required: true, message: "请上传视频", trigger: "blur" }
                ],
                imgs: [
                    { required: true, message: "请上传封面", trigger: "blur" }
                ]
            },
            topics: [],
            vSignal: new AbortController(),
            iSignal: new AbortController()
        }
    },
    computed: {
        ...mapState({
            categoryList: state => state.category.categoryList
        })
    },
    components: {
        VideoCut
    },
    methods: {
        getSensitive,
        initForm() {
            this.form = {
                desc: "",
                videos: [],
                imgs: [],
                topics: [],
                category: ''
            }
        },
        getTopic() {
            VideoTopicList().then(res => {
                if (res.status === 200) {
                    const topics = res.data.data
                    this.topics = topics || []
                }
            })
        },
        handleVideoChange(file, fileList) {
            if (!file.raw.type.includes("video")) {
                this.$message.error("视频格式不正确")
                fileList.splice(0, 1)
                return false
            }
            if (file.size > 20 * 1024 * 1024) {
                this.$message.error("视频大小不能超过 20M")
                fileList.splice(0, 1)
                return false
            }
            this.form.videos.splice(0, 1, Object.assign(file, {
                url: URL.createObjectURL(file.raw)
            }))
        },
        handleVideoRemove() {
            if (this.form.videos[0].url.includes("blob:")) {
                URL.revokeObjectURL(this.form.videos[0].url)
                this.vSignal.abort()
                this.vSignal = new AbortController()
            } else if (this.form.videos[0].url.includes("http")) {
                DeleteFile(this.form.videos[0].url)
            }
            this.form.videos.splice(0, 1)
        },
        handleVideoUpload({ file }) {
            if (!file.type.includes("video") || file.size > 20 * 1024 * 1024) {
                return false
            }
            const formData = new FormData();
            formData.append("file", file);
            UploadFile(formData, this.vSignal.signal).then(res => {
                if (!res) throw new Error("取消上传")
                if (res.status !== 200) throw new Error(res.data.msg);
                if (this.form.videos[0].uid === file.uid && this.form.videos[0].url.includes("blob:")) {
                    this.form.videos[0].url = res.data.data
                } else {
                    this.form.videos.splice(0, 1, Object.assign(file, {
                        url: res.data.data,
                        uid: file.uid,
                        name: file.name
                    }))
                }
            }).catch(err => {
                this.$message.error(err);
            })
        },
        handleImgChange(file, fileList) {
            this.form.imgs.splice(0, 1, Object.assign(file, {
                url: URL.createObjectURL(file.raw)
            }))
        },
        handleImgRemove() {
            if (this.form.imgs[0].url.includes("blob:")) {
                URL.revokeObjectURL(this.form.imgs[0].url)
                this.iSignal.abort()
                this.iSignal = new AbortController()
            } else if (this.form.imgs[0].url.includes("http")) {
                DeleteFile(this.form.imgs[0].url)
            }
            this.form.imgs.splice(0, 1)
        },
        getImg(raw) {
            this.form.imgs = [
                {
                    url: URL.createObjectURL(raw),
                    name: "截取封面",
                    raw
                }
            ]
            this.handleImgUpload({ file: raw })
        },
        handleImgUpload({ file }) {
            const formData = new FormData();
            formData.append("file", file);
            UploadFile(formData, this.iSignal.signal).then(res => {
                if (!res) throw new Error("取消上传")
                if (res.status !== 200) throw new Error(res.data.msg);
                if (this.form.imgs[0].uid === file.uid && this.form.imgs[0].url.includes("blob:")) {
                    URL.revokeObjectURL(this.form.imgs[0].url)
                    this.form.imgs[0].url = res.data.data
                } else {
                    this.form.imgs.splice(0, 1, Object.assign(file, {
                        url: res.data.data,
                        uid: file.uid
                    }))
                }
            }).catch(err => {
                this.$message.error(err);
            })
        },
        addTopic(e) {
            if (this.form.topics.find(_ => _ === e.target.value)) {
                this.form.topics.splice(this.form.topics.indexOf(e.target.value), 1)
            } else {
                if (this.form.topics.length >= 3) {
                    this.$message.warning("最多选择三个话题")
                    this.form.topics.splice(0, 1)
                }
            }
            this.form.topics.push(e.target.value)
        },
        submit() {
            const word = this.getSensitive(this.form.desc)
            if (word) {
                this.$message.error(`请勿发布敏感词「${word}」!`)
                return
            }
            if (this.isSubmit) return
            this.isSubmit = true
            this.$refs.form.validate((valid) => {
                if (!valid) {
                    this.isSubmit = false
                    return
                }
                const fn = () => {
                    if (this.form.videos[0].url.includes("blob:") || this.form.imgs[0].url.includes("blob:")) {
                        fn()
                        return
                    }
                    const form = {
                        categoryId: this.form.category,
                        description: this.form.desc,
                        topics: this.form.topics,
                        coverUrl: this.form.imgs[0].url,
                        playUrl: this.form.videos[0].url
                    }
                    VideoPublish(form).then(res => {
                        if (res.status === 200) {
                            this.$message.success(res.data.msg)
                            this.$refs.form.resetFields()
                            this.initForm()
                            setTimeout(() => {
                                this.$router.push({ name: "Home" })
                            }, 300)
                        } else {
                            this.$message.error(res.data.msg)
                        }
                    }).finally(() => {
                        this.isSubmit = false
                    })
                }

                setTimeout(() => {
                    fn()
                }, 1000 / 16)

            })
        }
    },
    mounted() {
        this.getTopic()
    }
};
</script>

<style lang="less" scoped>
.container {
    width: 600px;
    padding-bottom: 50px;
    margin: auto;

    .title {
        font-size: 24px;
        line-height: 28px;
        color: @white;
        margin: 20px 0;
    }

    .upload-container {
        height: fit-content !important;

        .el-form-item {

            /deep/ .el-form-item__label {
                width: 100%;
                font-size: 16px;
                color: @white;
                display: flex;
                float: unset;
                align-items: flex-end;

                .count {
                    font-size: 12px;
                    font-style: normal;
                    margin-left: auto;
                }

                .cut {
                    font-size: 12px;
                    font-style: normal;
                    padding: 0 16px;
                    height: 32px;
                    background-color: @red;
                    color: @white;
                    margin-left: auto;
                    border: unset;
                }
            }

        }



        .upload {
            height: 240px;
            width: 600px;

            &.upload-pic {
                width: 360px;
                height: 200px;
            }

            /deep/ .el-upload {
                height: calc(100% - 25px);
                width: 100%;
                color: @gray-2;

                .el-upload-dragger {
                    height: 100%;
                    width: 100%;
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                    justify-content: center;
                    background-color: @transparent-dark;
                    border: unset;

                    .el-icon-upload {
                        margin-top: 0;
                    }

                    em {
                        color: @red;
                    }
                }


            }

            /deep/ .el-upload-list {

                .el-upload-list__item {
                    cursor: default;

                    &:hover {
                        background-color: @gray-4;
                    }

                    .el-upload-list__item-name {
                        &:hover {
                            color: @gray-2;
                        }
                    }
                }
            }
        }

        .el-textarea,
        .el-input,
        .el-select {
            background-color: @transparent-dark;
            font-size: 14px;
            border-radius: 4px;
            line-height: 20px;
            color: @gray-1;
            width: 100%;

            /deep/ .el-select__input {
                margin-left: 10px;
            }

            /deep/ .el-textarea__inner,
            /deep/ .el-input__inner {
                outline: none;
                border: none;
                background-color: transparent;
                height: 100%;
                color: @gray-1;
                padding: 10px;
                font-family: inherit;

                &::-webkit-scrollbar {
                    display: none;
                }

                &::placeholder {
                    color: @gray-2;
                }
            }

            /deep/ .el-input__icon {
                line-height: inherit;
            }
        }


        .submit {
            width: 100%;
            margin-top: 20px;
            background-color: @red;
            color: @white;
            font-size: 16px;
            line-height: 20px;
            border: unset;
        }

        .el-button {
            &:hover {
                filter: brightness(.9);
            }

            &:active {
                filter: brightness(.8);
            }

            &:disabled {
                filter: brightness(.3) !important;
            }
        }
    }
}
</style>
