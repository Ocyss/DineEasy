<template>
  <view>
    <SelectStoreVue v-if="!selectFlagRef" />
    <view v-else>
      <view class="search-input">
        <uni-search-bar
          @input="inputend"
          cancelButton="none"
          clearButton="auto"
          placeholder="请输入关键词"
          :radius="100" />
      </view>
      <view class="store-data">
        <view
          class="left"
          @click="selectFlagRef = false">
          <view class="name">{{ storeDataRef.name }}</view>
          <view class="address">{{ storeDataRef.address }}</view>
        </view>
        <view class="typeview">
          <button
            class="type"
            @click="changeType">
            <text v-if="DiningStyleRef == 1">堂食</text>
            <text v-if="DiningStyleRef == 2">自提</text>
            <text v-if="DiningStyleRef == 3">外卖</text>
          </button>
        </view>
      </view>
      <view class="content">
        <MenuVue />
      </view>
      <ChosenDishesVue />
      <uni-popup
        ref="popup"
        :animation="false"
        background-color="#fff">
        <view class="popup-content">
          <view class="popup-button">
            <button
              type="primary"
              plain="true"
              @click="changeDining(1)"
              :disabled="DiningStyleRef == 1">
              堂食
            </button>
            <button
              type="primary"
              plain="true"
              @click="changeDining(2)"
              :disabled="DiningStyleRef == 2">
              自提
            </button>
            <button
              type="primary"
              plain="true"
              @click="changeDining(3)"
              :disabled="DiningStyleRef == 3">
              外卖
            </button>
          </view>
        </view>
      </uni-popup>
    </view>
  </view>
</template>

<script setup>
  import { onLoad, onShow } from '@dcloudio/uni-app';
  import MenuVue from '@/components/order/menu/menu.vue';
  import ChosenDishesVue from '@/components/order/menu/ChosenDishes.vue';
  import SelectStoreVue from '@/components/order/selectStore/selectStore.vue';
  import { ref } from 'vue';
  const storeDataRef = ref(uni.getStorageSync('Store'));
  const DiningStyleRef = ref(
    uni.getStorageSync('DiningStyle') ? uni.getStorageSync('DiningStyle') : 1
  );
  const selectFlagRef = ref(storeDataRef.value.id != null);
  const popup = ref();
  uni.$on('selectOK', function () {
    storeDataRef.value = uni.getStorageSync('Store');
    selectFlagRef.value = true;
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

  function changeType() {
    popup.value.open('bottom');
  }
  function changeDining(id) {
    uni.setStorageSync('DiningStyle', id);
    DiningStyleRef.value = id;
    popup.value.close();
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

  .popup-content {
    width: 100%;
    position: absolute;
    bottom: var(--window-bottom);
    background-color: #fff;
    display: flex;
    justify-content: center;
    align-items: center;
    .popup-button {
      width: 80%;
      :deep(uni-button) {
        margin: 20px;
      }
    }
  }
</style>
