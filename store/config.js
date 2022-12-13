import {
    defineStore
} from 'pinia';

export const useConfig = defineStore('config', { // 定义要状态数据
    state: () => {
        return {
            Store: {
                id: null, //店铺id
                name: null, //店铺名
                clicked: null, //购物车
                tableId: null, //桌号
                type: null, //堂食还是自提，外卖
                address: null,
            }, //当前选择的店铺
            Mobile: '', //手机号
            UserName: '邱爹', //用户名
        };
    },
    actions: {},
});
