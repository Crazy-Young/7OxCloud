const mobileReg = /^1[3456789]\d{9}$/;
const captchaReg = /^\d{6}$/;
const passwordReg = /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,20}$/;

export {
    mobileReg,
    captchaReg,
    passwordReg
}