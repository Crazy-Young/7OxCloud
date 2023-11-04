<template>
    <div id="change-info">
        <!-- 遮罩层 -->
        <div class="mask" @click="$emit('close')"></div>
        <!-- 修改面板 -->
        <div class="panel">
            <!-- 关闭按钮 -->
            <div class="panel-close" @click="$emit('close')"></div>
            <!-- 修改面板头部 -->
            <div class="panel__header">
                <div class="panel__header-title">编辑资料</div>
            </div>
            <!-- 修改面板内容 -->
            <div class="panel-content">
                <!-- 修改表单 -->
                <el-form :model="form" ref="form" :show-message="false" label-position="top">
                    <el-form-item prop="avatar">
                        <label class="avatar" for="avatar">
                            <img :src="form.avatar" alt="">
                            <div class="avatar-mask"></div>
                        </label>
                        <input type="file" accept="image/*" name="avatar" id="avatar" @change="changeAvatar">
                    </el-form-item>

                    <el-form-item prop="username" label="昵称">
                        <el-input v-model="form.username" :placeholder="userInfo.username" maxlength="20">
                            <template slot="append">
                                <span class="count">{{ form.username.length }}/20</span>
                            </template>
                        </el-input>
                    </el-form-item>

                    <el-form-item prop="age" label="年龄">
                        <el-input v-model.number="form.age" @input="limitAge" :placeholder="userInfo.age" type="number"
                            :min="0" :step="1" :max="100">
                        </el-input>
                    </el-form-item>

                    <el-form-item prop="gender" label="性别">
                        <el-select v-model="form.gender" placeholder="选择性别">
                            <el-option label="男" value="1"></el-option>
                            <el-option label="女" value="0"></el-option>
                        </el-select>
                    </el-form-item>

                    <el-form-item prop="location" label="地区">
                        <el-input v-model="form.location" :placeholder="userInfo.location" maxlength="20">
                        </el-input>
                    </el-form-item>

                    <el-form-item prop="signature" label="简介">
                        <el-input v-model="form.signature" :placeholder="userInfo.signature" maxlength="200" type="textarea"
                            resize="none" :rows="4">
                        </el-input>
                    </el-form-item>

                    <el-form-item>
                        <el-button type="primary" @click="submitForm('form')">保存</el-button>
                        <el-button type="info" @click="$emit('close')">取消</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </div>
</template>

<script>
import { mapState } from "vuex";
import { UploadFile, DeleteFile } from "@/api/base";
export default {
    name: "ChangeInfo",
    data() {
        return {
            form: {
                username: '',
                signature: '',
                avatar: '',
                location: '',
                gender: '',
                age: 0,
            },
            file: null
        };
    },
    methods: {
        limitAge() {
            if (this.form.age > 100) {
                this.form.age = 100
            } else if (this.form.age < 0) {
                this.form.age = 0
            }
        },
        submitForm(formName) {
            this.$refs[formName].validate((valid, error) => {
                if (valid) {
                    // 如果有文件，上传后再修改
                    if (this.file) {
                        const formData = new FormData();
                        formData.append("file", this.file);
                        UploadFile(formData).then(res => {
                            if (res.status !== 200) throw new Error(res.data.msg);
                            // 删除原来的头像
                            DeleteFile(this.userInfo.avatar)
                            // 修改头像url为上传后的
                            this.form.avatar = res.data.data;
                            // 修改用户信息
                            return this.$store.dispatch("UserUpdate", {
                                username: this.form.username || this.userInfo.username,
                                signature: this.form.signature || this.userInfo.signature,
                                avatar: this.form.avatar || this.userInfo.avatar,
                                location: this.form.location || this.userInfo.location,
                                gender: this.form.gender || this.userInfo.gender,
                                age: this.form.age || this.userInfo.age
                            })
                        }).then((res) => {
                            // 关闭弹窗
                            this.$emit("close");
                        }).catch(err => {
                            // 错误: 打印错误信息
                            this.$message.error(err);
                        })
                    } else {
                        // 修改用户信息
                        this.$store.dispatch("UserUpdate", {
                            username: this.form.username || this.userInfo.username,
                            signature: this.form.signature || this.userInfo.signature,
                            avatar: this.form.avatar || this.userInfo.avatar,
                            location: this.form.location || this.userInfo.location,
                            gender: this.form.gender.toString() || this.userInfo.gender.toString(),
                            age: this.form.age || this.userInfo.age
                        }).then((res) => {
                            // 关闭弹窗
                            this.$emit("close");
                        })
                    }
                }
            });
        },
        changeAvatar(e) {
            const file = e.target.files[0];
            this.file = file;
            this.form.avatar = URL.createObjectURL(file);
        }
    },
    computed: {
        ...mapState({
            userInfo: state => state.user.userInfo,
        }),
    },
    mounted() {
        // 初始化表单
        this.form.username = this.userInfo.username;
        this.form.signature = this.userInfo.signature;
        this.form.avatar = this.userInfo.avatar;
        this.form.location = this.userInfo.location;
        this.form.gender = this.userInfo.gender.toString();
        this.form.age = this.userInfo.age;
    }
};
</script>

<style lang="less" scoped>
#change-info {
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
        min-height: 588px;
        padding: 0 40px;
        flex-direction: column;
        transform: translate(-50%, -50%);
        width: 480px;
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

            .panel__header-sub {
                font-size: 12px;
                font-weight: 400;
                height: 22px;
                line-height: 22px;
                margin-top: 12px;
            }

        }

        .panel-content {
            flex: 1;
            border-radius: 12px;
            min-height: 320px;
            display: flex;

            .el-form {
                width: 100%;

                .el-form-item {
                    width: 100%;
                    margin-bottom: 10px;

                    /deep/ .el-form-item__label {
                        color: @white;
                        line-height: 2;
                        padding-bottom: 0;
                    }

                    &:last-child {
                        margin-top: 20px;
                        margin-bottom: 0;
                    }

                    .avatar {
                        height: 108px;
                        width: 108px;
                        border-radius: 50%;
                        position: relative;
                        overflow: hidden;
                        display: block;
                        margin: 16px auto;

                        &~#avatar {
                            display: none;
                        }

                        img {
                            height: 100%;
                            width: 100%;
                            object-fit: cover;
                            border-radius: 50%;
                        }

                        &-mask {
                            position: absolute;
                            top: 0;
                            left: 0;
                            height: 100%;
                            width: 100%;
                            display: block;
                            z-index: 1;
                            background-color: rgba(0, 0, 0, .3);
                            background-position: 50%;
                            background-repeat: no-repeat;
                            background-size: 32px 32px;
                            background-image: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMzIiIGhlaWdodD0iMzIiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGQ9Ik0xMi41MzUgNWEyIDIgMCAwIDAtMS42NjQuODkxTDkuNDY0IDguMDAzSDhhNCA0IDAgMCAwLTQgNFYyM2E0IDQgMCAwIDAgNCA0aDE2YTQgNCAwIDAgMCA0LTRWMTIuMDAzYTQgNCAwIDAgMC00LTRoLTEuNDY1TDIxLjEzIDUuODkxQTIgMiAwIDAgMCAxOS40NjQgNWgtNi45MjlaTTE2IDEwLjVBNi4yNSA2LjI1IDAgMSAwIDE2IDIzYTYuMjUgNi4yNSAwIDAgMCAwLTEyLjVabS0zLjc1IDYuMjVhMy43NSAzLjc1IDAgMSAxIDcuNSAwIDMuNzUgMy43NSAwIDAgMS03LjUgMFoiIGZpbGw9IiNmZmYiIG9wYWNpdHk9Ii45Ii8+PC9zdmc+);
                        }
                    }
                }
            }

            .el-input {
                background-color: @gray-3;
                font-size: 14px;
                height: 32px;
                line-height: 32px;
                border-radius: 4px;
                color: @gray-1;

                /deep/ .el-input__inner {
                    outline: none;
                    border: none;
                    background-color: transparent;
                    height: 100%;
                    color: @gray-1;
                    padding-left: 10px;
                    font-family: inherit;

                    &::placeholder {
                        color: @gray-2;
                    }



                    &::-webkit-inner-spin-button,
                    &::-webkit-outer-spin-button {
                        -webkit-appearance: none;
                    }
                }

                /deep/ .el-input-group__append {
                    border: unset;
                    background-color: transparent;
                    cursor: pointer;
                    padding: 0 10px;

                    .count {
                        color: @gray-2;
                        font-size: 12px;
                    }
                }
            }

            .el-select {
                background-color: @gray-3;
                font-size: 14px;
                height: 32px;
                line-height: 32px;
                border-radius: 4px;
                color: @gray-1;

                /deep/ .el-input__inner {
                    outline: none;
                    border: none;
                    background-color: transparent;
                    height: 100%;
                    color: @gray-1;
                    padding-left: 10px;
                    font-family: inherit;

                    &::placeholder {
                        color: @gray-2;
                    }
                }

                /deep/ .el-input__suffix {

                    .el-input__icon {
                        line-height: inherit;
                    }
                }

                /deep/ .el-input-group__append {
                    border: unset;
                    background-color: transparent;
                    cursor: pointer;
                    padding: 0 10px;

                    .count {
                        color: @gray-2;
                        font-size: 12px;
                    }
                }
            }

            .el-textarea {
                background-color: @gray-3;
                font-size: 14px;
                border-radius: 4px;
                line-height: 20px;
                color: @gray-1;

                /deep/ .el-textarea__inner {
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
            }

            .el-button {
                background-color: @red;
                border: unset;
                padding: 0 16px;
                height: 36px;
                line-height: 22px;
                min-width: 88px;
                font-size: 14px;

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