<template>
  <view>
    <SelectStoreVue v-if="!select_flag" />
    <view v-else>
      <view class="search-input">
        <uni-search-bar
          @input="inputend"
          cancelButton="none"
          clearButton="auto"
          placeholder="请输入关键词"
          :radius="100" />
      </view>
      <view
        class="store-data"
        @click="select_flag = false">
        <view class="left">
          <view class="name">{{ config.Store.name }}</view>
          <view class="address">{{ config.Store.address }}</view>
        </view>
        <view class="typeview"><button class="type">自提</button></view>
      </view>
      <view class="content">
        <MenuVue />
      </view>
      <ChosenDishesVue />
    </view>
  </view>
</template>

<script setup>
  import { useConfig } from '@/store/config.js';
  import { onLoad, onShow } from '@dcloudio/uni-app';
  import MenuVue from '@/components/order/menu/menu.vue';
  import ChosenDishesVue from '@/components/order/menu/ChosenDishes.vue';
  import SelectStoreVue from '@/components/order/selectStore/selectStore.vue';
  import { ref } from 'vue';
  const config = useConfig();
  const select_flag = ref(false);
  uni.$on('selectOK', function () {
    select_flag.value = true;
  });
  let inputTime = null;

  //搜索单品
  function search(val) {
    console.log(val);
  }

  //输入完成，采用防抖
  function inputend(val) {
    if (inputTime != null) {
      clearTimeout(inputTime);
    }
    inputTime = setTimeout(() => {
      search({
        value: val
      });
    }, 600);
  }
</script>

<style scoped lang="scss">
  @import '@/uni.scss';

  // 搜索框
  .search-input {
    height: 8vh;
  }

  // 门店信息
  .store-data {
    border-bottom: 1px solid #b3b3b3;
    display: flex;
    justify-content: space-between;
    height: 10vh;
    padding: 0px 15px;
    box-sizing: border-box;

    // 左边部分
    .left {
      display: flex;
      flex-direction: column;
      justify-content: space-evenly;
    }

    // 门店名称
    .name {
      font-size: 1.3rem;
      font-weight: 600;
      &::after {
        content: '﹥';
      }
    }

    // 门店地址
    .address {
      font-size: 0.8rem;
      color: #797979;
    }
    //外卖类型
    .typeview {
      display: flex;
      justify-content: center;
      align-items: center;
      .type {
        display: flex;
        justify-content: center;
        align-items: center;
        background-color: $theme-color;
        color: #fff;
        width: 65px;
        height: 27px;
        border-radius: 20px;
      }
    }
  }
</style>
