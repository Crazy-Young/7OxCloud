<template>
    <div id="login">
        <!-- 遮罩层 -->
        <div class="mask" @click="$emit('close')"></div>
        <!-- 登陆面板 -->
        <div class="panel">
            <!-- 关闭按钮 -->
            <div class="panel-close" @click="$emit('close')"></div>
            <!-- 面板顶部 -->
            <div class="panel__header">
                <!-- 标题: 根据isReset判断不同顶部信息 -->
                <div class="panel__header-title">{{ isReset ? '重置密码' : '登录后体验更多功能' }}</div>
                <div class="panel__header-sub" v-if="!isReset">
                    <ul class="panel__header-sub-list">
                        <li id="comment">点赞评论随心发</li>
                        <li id="friend">朋友视频一键览</li>
                    </ul>
                </div>
            </div>
            <!-- 面板内容: 登录 -->
            <div class="panel-content" v-if="!isReset">
                <el-tabs v-model="activeChoice">
                    <el-tab-pane label="密码登录" name="password">
                        <el-form @submit="submitForm('password')" :model="form" :rules="rules" ref="password"
                            :show-message="false">
                            <el-form-item prop="mobile">
                                <el-input v-model="form.mobile" placeholder="手机号" maxlength="11"></el-input>
                            </el-form-item>
                            <el-form-item prop="password">
                                <el-input type="password" show-password placeholder="密码" maxlength="20" minlength="8"
                                    v-model="form.password" autocomplete="off">
                                    <template #append>
                                        <button class="append" id="forget" type="button"
                                            @click="isReset = true">忘记密码</button>
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="captcha">
                                <el-input v-model="form.captcha" placeholder="验证码" maxlength="6">
                                    <template #append>
                                        <img :src="captchaUrl" class="append" @click="SendCaptcha" />
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item>
                                <el-button type="primary" @click="submitForm('password')">登录</el-button>
                            </el-form-item>
                        </el-form>
                    </el-tab-pane>
                    <el-tab-pane label="验证码登录" name="captcha">
                        <el-form @submit="submitForm('captcha')" :model="form" :rules="rules" ref="captcha"
                            :show-message="false">
                            <el-form-item prop="mobile">
                                <el-input v-model="form.mobile" placeholder="手机号" maxlength="11"></el-input>
                            </el-form-item>
                            <el-form-item prop="captcha">
                                <el-input v-model="form.captcha" placeholder="验证码" maxlength="6">
                                    <template #append>
                                        <button class="append" type="button" @click="sendSMS(2)"
                                            :disabled="!validMobile || seconds > 0">{{ seconds > 0 ? `重新发送(${seconds}s)` :
                                                '获取验证码' }}</button>
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item>
                                <p class="tips">如未收到验证码请检查手机号是否填写正确</p>
                                <el-button type="primary" @click="submitForm('captcha')">登录</el-button>
                            </el-form-item>
                        </el-form>
                    </el-tab-pane>
                    <el-tab-pane label="注册" name="register">
                        <el-form @submit="submitForm('register')" :model="form" :rules="rules" ref="register"
                            :show-message="false">
                            <el-form-item prop="mobile">
                                <el-input v-model="form.mobile" placeholder="手机号" maxlength="11"></el-input>
                            </el-form-item>
                            <el-form-item prop="captcha">
                                <el-input v-model="form.captcha" placeholder="验证码" maxlength="6">
                                    <template #append>
                                        <button class="append" type="button" @click="sendSMS(1)"
                                            :disabled="!validMobile || seconds > 0">{{ seconds > 0 ? `重新发送(${seconds}s)` :
                                                '获取验证码' }}</button>
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="password">
                                <el-input v-model="form.password" placeholder="密码" maxlength="20" minlength="8"
                                    type="password" show-password>
                                </el-input>
                            </el-form-item>
                            <el-form-item>
                                <el-button type="primary" @click="submitForm('register')">登录</el-button>
                            </el-form-item>
                        </el-form>
                    </el-tab-pane>
                </el-tabs>
            </div>

            <!-- 面板内容: 重置 -->
            <div class="panel-content reset" v-else>
                <el-form :model="form" :rules="rules" ref="reset" :show-message="false">
                    <el-form-item prop="mobile">
                        <el-input v-model="form.mobile" placeholder="手机号" maxlength="11"></el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <el-input type="password" show-password placeholder="修改密码" maxlength="20" minlength="8"
                            v-model="form.password" autocomplete="off">
                        </el-input>
                    </el-form-item>
                    <el-form-item prop="captcha">
                        <el-input v-model="form.captcha" placeholder="验证码" maxlength="6">
                            <template #append>
                                <button class="append" type="button" @click="sendSMS(3)"
                                    :disabled="!validMobile || seconds > 0">{{ seconds > 0 ? `重新发送(${seconds}s)` :
                                        '获取验证码' }}</button>
                            </template>
                        </el-input>
                    </el-form-item>
                    <el-form-item>
                        <p class="tips" @click="isReset = false"><i class="el-icon-back icon"></i>取消重置</p>
                        <el-button type="primary" @click="submitForm('reset')">修改</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </div>
</template>

<script>
import { mobileReg, captchaReg, passwordReg } from '@/utils/validate'
import { SendCaptcha, SendSMS } from '@/api/base';
export default {
    name: "Login",
    data() {
        return {
            activeChoice: 'password',
            form: {
                mobile: '',
                password: '',
                captcha: ''
            },
            isReset: false,
            captchaId: '',
            captchaUrl: '',
            timer: null,
            seconds: -1
        };
    },
    methods: {
        submitForm(formName) {
            this.$refs[formName].validate((valid, error) => {
                if (valid) {
                    if (this.isReset) {
                        this.passwordReset();
                        return
                    }
                    switch (this.activeChoice) {
                        case 'password':
                            this.login(1);
                            break;
                        case 'captcha':
                            this.login(2);
                            break;
                        case 'register':
                            this.register();
                            break
                    }
                } else {
                    this.$message.error(Object.values(error)[0][0].message);
                    return false;
                }
            });
        },
        resetForm(formName) {
            this.$refs[formName].resetFields();
        },
        sendSMS(type) {
            this.seconds = 60;
            SendSMS({
                type,
                mobile: this.form.mobile
            }).then(res => {
                if (res?.data?.code === 200) {
                    this.timer = setInterval(() => {
                        this.seconds--;
                        if (this.seconds <= 0) {
                            clearInterval(this.timer);
                            this.seconds = -1;
                            this.timer = null;
                        }
                    }, 1000)
                    this.$message.success(res?.data?.msg);
                } else {
                    this.seconds = -1;
                    this.$message.error(res?.data?.msg);
                }
            })
        },
        SendCaptcha() {
            SendCaptcha().then(res => {
                if (res?.data?.code === 200) {
                    this.captchaId = res?.data?.data.captchaId;
                    this.captchaUrl = res?.data?.data.picPath;
                } else {
                    this.$message.error(res?.data?.msg);
                }
            })
        },
        register() {
            this.$store.dispatch("UserRegister", {
                mobile: this.form.mobile,
                password: this.form.password,
                captcha: this.form.captcha
            }).then(res => {
                this.$emit('close');
                location.reload();
            })
        },
        login(type) {
            this.$store.dispatch("UserLogin", {
                type,
                mobile: this.form.mobile,
                password: this.form.password,
                captcha: this.form.captcha,
                captchaId: this.captchaId
            }).then(res => {
                this.$emit('close');
                location.reload();
            }).catch(err => {
                this.SendCaptcha();
            })
        },
        passwordReset() {
            this.$store.dispatch("UserPasswordReset", {
                mobile: this.form.mobile,
                password: this.form.password,
                captcha: this.form.captcha
            }).then(res => {
                this.isReset = false;
                this.resetForm('reset');
            })
        }
    },
    computed: {
        rules() {
            return {
                mobile: [
                    { required: true, message: '请输入手机号', trigger: 'blur' },
                    { min: 11, max: 11, message: '手机号长度为 11 个字符', trigger: 'blur' },
                    { pattern: mobileReg, message: '请输入正确的手机号', trigger: 'blur' },
                ],
                password: [
                    { required: this.activeChoice !== 'captcha', message: '请输入密码', trigger: 'blur' },
                    { min: 8, max: 20, message: '密码长度在 8 到 20 个字符', trigger: 'blur' },
                    { pattern: passwordReg, message: '密码必须同时包含数字和字母', trigger: 'blur' },
                ],
                captcha: [
                    { required: true, message: '请输入验证码', trigger: 'blur' },
                    { min: 6, max: 6, message: '验证码长度为 6 个字符', trigger: 'blur' },
                    { pattern: captchaReg, message: '请输入正确的验证码', trigger: 'blur' },
                ]
            }
        },
        validMobile() {
            return mobileReg.test(this.form.mobile)
        }
    },
    created() {
        this.SendCaptcha();
    }
};
</script>

<style lang="less" scoped>
#login {
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
        flex-direction: column;
        transform: translate(-50%, -45%);
        width: 360px;
        min-height: 450px;
        padding: 18px 28.5px 30px;
        background: url('../assets/login_bg.png') 0% 0% / 100% no-repeat rgb(255, 255, 255);
        border-radius: 8px;

        .panel-close {
            background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAJAAAACQCAYAAADnRuK4AAAACXBIWXMAACxLAAAsSwGlPZapAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAPTSURBVHgB7d29ThRRGIDh7wyLWGlFQoJ7A4rBBgs7CxPs8QqwsRMra3u8Am/AYGOnnR2yGxNZxM4KsgnQURrY43y7IQgL7PzPnHPeJzFCmGLIvMz/OSsCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoCxGPLGwsNSemjJrVuxzY+SOtWZPrN0UidZ7va19aQBdxygyqyaSZRHbtlaOxcqXJq1jWl4E9PDR4xfG2ncazuWfaUgnYld/b3d3pUaLi0sPrMinq9ZRDcS83vnZ2RDHOB/QcM/Tku83LaN/6SciK3VFNCkeNdobRc9c2xNF4rgokjeTltEN14o34P14Q0rFksSj9OfG2Im/S9M4H1C8D11OtFgNESWN54yev4ljnA8o6cY5W7aqiNLGo9Is2xTu74FEUp0zVBFRlnhG4itHx7gf0PAyOJ0yI8oez/AQtimOcT6g01vyYXgFk1IZEeWJZ7j3mYnWxTHOB/Sr292L70Ws1B1RrnhMvO7TdrW35d7NRG/uROc6dOS8T5Q7npasbHfrvdGZlTcBqToiCjke5VVAqsqIQo9HeReQqiIi4hnxMiBVZkTEc87bgFQZERHPRV4HpIqMiHjGeR+QKiKiaf2aeMYEEZDKG5H+TzzjgglI5XvUkIHn8aigAlKVRRRAPCq4gFTpEQUSjwoyIFVaRAHFo4INSBUeUWDxqKADUoVFFGA8KviAVO6IAo1HtQS52YHIyV8JEoewgg5hdQ9erAsn0QWeRIcYEZfxBV/GhxYRNxJLEFJEPMooSSgR8TC1RCFE5MPQ5kTyvgzWlMGLTRNEQEW8SdiEwYtNxCutN7l0h7nOwYtNxUv117nm8QQRXcSwnqtMeLZFROcYWHhZwgejRDTC0Ob/pXyqTkRMrnAu4ysZoUfE9C4q5/s8IUfkRUBNGDEaakReTDQeTQ033D1Jq+A3CXMPXmSi8erpRONNiEdtx3uQPHesJbIvxTE+TDT+RNIq8R3mXBHZZJOmN4kPE42n2/tU8AJ85oiMbYtjnA8o1UaqcPREloiy7LXq5sMhLNlE4zUMvUkdkUk/aXrdnA9oMC3vJ26gGsdtpYrIwYnGp8Rxh/3+8ezc/FF03af2NGDQ38FB/2hubv5bvC5P42/HL/HjdTy18nbnR8e5jzpwPiB1eNDfnW3Pb5iBfuaW6InoTPxv34r5GP9Vr/U6nT9SM40oXsevw3UUuSsakr7pOJDPcjt65WI8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQJX+ATc9oF5n57RZAAAAAElFTkSuQmCC);
            background-size: 100% 100%;
            cursor: pointer;
            height: 27px;
            position: absolute;
            right: 10px;
            top: 10px;
            width: 27px;
            z-index: 9999;
        }

        .panel__header {
            background: transparent;
            border-radius: 4px 4px 0 0;
            margin: 0 auto 12px;
            text-align: center;
            width: 100%;

            .panel__header-title {
                font-size: 18px;
                font-weight: 500;
                height: 34px;
                line-height: 34px;
                color: @black;
            }

            .panel__header-sub {
                font-size: 12px;
                font-weight: 400;
                height: 22px;
                line-height: 22px;
                margin-top: 12px;

                .panel__header-sub-list {
                    list-style: none;
                    display: flex;
                    gap: 20px;
                    justify-content: center;
                    line-height: 1;

                    #comment,
                    #friend {
                        cursor: pointer;
                        position: relative;
                        padding-left: 20px;
                    }

                    #comment::before,
                    #friend::before {
                        content: "";
                        height: 18px;
                        left: 0;
                        position: absolute;
                        top: 50%;
                        transform: translateY(-50%);
                        width: 18px;
                    }

                    #comment::before {
                        background: url(data:image/png;base64,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) no-repeat 50%/contain;
                    }

                    #friend::before {
                        background: url(data:image/png;base64,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) no-repeat 50%/contain;
                    }
                }
            }


        }

        .panel-content {
            flex: 1;
            border-radius: 12px;
            box-shadow: 0 0 30px -22px @red;
            min-height: 320px;
            display: flex;
            background-color: rgba(255, 255, 255, 0.2);

            &.reset {
                display: flex;
                align-items: center;
                padding: 36px;

                .tips {
                    cursor: pointer;
                    width: fit-content;
                }
            }

            .el-tabs {
                flex: 1;
                display: flex;
                flex-direction: column;

                /deep/ .el-tabs__header {
                    padding: 0 25px;
                    margin: 0;

                    .el-tabs__item {
                        line-height: 50px;
                        height: 50px;
                        color: @gray-2;

                        &.is-active {
                            color: @black;
                        }
                    }

                    .el-tabs__active-bar {
                        background-color: @red;
                        transition: all .3s ease-in-out;
                    }

                    .el-tabs__nav-wrap::after {
                        height: 1px;
                        transform: scaleY(.5);
                    }

                    .el-tabs__nav-scroll {
                        display: flex;
                        justify-content: center;
                    }
                }

                /deep/ .el-tabs__content {
                    padding: 0 36px;
                    flex: 1;
                    display: flex;
                    flex-direction: column;
                    justify-content: center;
                }
            }

            .el-form {
                .el-form-item {

                    &:last-child {
                        margin-top: 20px;
                        margin-bottom: 0;
                    }

                    .tips {
                        font-size: 12px;
                        line-height: 3;
                        transform: scale(.9);
                        transform-origin: left;
                        color: @gray-2;

                        .icon {
                            margin-right: 5px;
                        }
                    }
                }
            }

            .el-input {
                background-color: @gray-0;
                border-radius: 4px;
                font-size: 12px;

                /deep/ .el-input__inner {
                    outline: none;
                    border: none;
                    background-color: transparent;
                }

                &.el-input-group {
                    display: flex;

                    /deep/ .el-input-group__append {
                        border: unset;
                        background-color: transparent;
                        cursor: pointer;
                        display: flex;
                        width: fit-content;
                        color: @red;

                        &:has(#forget) {
                            padding-left: 5px;
                        }

                        &:has(img) {
                            padding: 0;
                        }

                        .append {
                            all: unset;
                            max-height: 40px;

                            &:disabled {
                                color: @gray-2;
                                cursor: not-allowed;
                            }
                        }
                    }
                }

            }

            .el-button {
                width: 100%;
                background-color: @red;
                border: unset;

                &:hover {
                    filter: brightness(0.9);
                }

                &:active {
                    filter: brightness(0.8);
                }
            }
        }
    }


}
</style>