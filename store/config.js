import {
  defineStore
} from 'pinia';

export const useConfig = defineStore('config', { // 定义要状态数据
  state: () => {
    return {
      Mobile: '', //手机号
      UserName: '邱爹', //用户名
    };
  },
  actions: {},
});
