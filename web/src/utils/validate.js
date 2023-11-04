/**
 * @description: 校验手机号码: 11位
 * @author: 涂国彬
 */
const mobileReg = /^1[3456789]\d{9}$/;

/**
 * @description: 校验验证码: 6位
 * @author: 涂国彬
 */
const captchaReg = /^\d{6}$/;

/**
 * @description: 校验密码: 至少包含大小写字母和数字的组合
 * @author: 涂国彬
 */
const passwordReg = /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,20}$/;

export {
    mobileReg,
    captchaReg,
    passwordReg
}