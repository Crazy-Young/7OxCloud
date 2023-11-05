/**
 * @description: 防抖函数
 * @param {Function} fn 函数
 * @param {Number} delay 延迟时间
 * @param {Boolean} immediate 是否立即执行
 * @returns {Function}
 * @author: 涂国彬
 * @date: 2023-11-03
 */

export default function debounce(fn, delay = 300, immediate = false) {
    let timer = null
    let isInvoke = false

    const _debounce = function (...args) {
        if (timer) clearTimeout(timer)

        if (immediate && !isInvoke) {
            fn.apply(this, args)
            isInvoke = true
        } else {
            timer = setTimeout(() => {
                fn.apply(this, args)
                isInvoke = false
            }, delay)
        }
    }

    return _debounce
}
