/**
 * @description: 防抖函数
 * @param {Function} func 函数
 * @param {Number} wait 延迟时间
 * @param {Boolean} immediate 是否立即执行
 * @returns {Function}
 * @author: 涂国彬
 * @date: 2023-11-03
 */

export default function debounce(func, wait = 300, immediate = false) {
    let timeout;
    return function () {
        const context = this;
        const args = arguments;
        const later = function () {
            timeout = null;
            if (!immediate) func.apply(context, args);
        };
        const callNow = immediate && !timeout;
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
        if (callNow) func.apply(context, args);
    };
}