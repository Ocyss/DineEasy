//封装 公共 节流器
export const throttle = (fn, delay = 200) => {
    let timer = null
    return function() {
        if (timer) return
        timer = setTimeout(() => {
            fn.apply(this, arguments)
            timer = null
        }, delay)
    }
}
//封装 公共 防抖器
export const debounce = (fn, delay = 200) => {
    let time = null;
    return function() {
        if (time !== null) {
            clearTimeout(time);
        }
        time = setTimeout(() => {
            fn.call(this);
        }, delay)
    }
}
