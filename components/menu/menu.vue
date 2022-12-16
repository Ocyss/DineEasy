<template>
  <view class="container">
    <scroll-view
      enable-flex
      scroll-y
      scroll-with-animation
      class="menu">
      <view
        class="item"
        :class="{
          active: menuScrollRef === item.id || (menuIdRef === menuScrollRef) === item.id
        }"
        v-for="item in dishesRef"
        :key="item.id"
        @click="menuIdRef = item.id">
        <image
          class="menuImg"
          :src="item.picture"
          mode="widthFix"></image>
        {{ item.name }}
      </view>
    </scroll-view>
    <scroll-view
      class="sub-menu"
      scroll-y
      :scroll-into-view="'tab-' + menuIdRef"
      scroll-with-animation
      show-scrollbar="true"
      @scroll="scroll">
      <view
        ref="tabRef"
        class="menuItem"
        v-for="menuItem in dishesRef"
        :id="'tab-' + menuItem.id"
        :key="menuItem.id">
        <view class="title">
          <image
            class="menuImg"
            :src="menuItem.picture"
            mode="widthFix"></image>
          {{ menuItem.name }}
        </view>
        <view
          class="dishesItem"
          v-for="dishesItem in menuItem.items"
          :key="dishesItem.id">
          <image
            class="dishesImage"
            :src="dishesItem.picture"
            mode="widthFix"></image>
          <view class="info">
            <view class="name">{{ dishesItem.name }}</view>
            <view class="bo">
              <view class="jiage">
                ￥{{ dishesItem.price }}
                <span class="jiage-qi">起</span>
              </view>
              <button class="jiagou">加购</button>
            </view>
          </view>
        </view>
      </view>
    </scroll-view>
  </view>
</template>

<script setup>
  import { ref } from 'vue';
  import { onLoad, onReady } from '@dcloudio/uni-app';
  import api from '@/api';
  const menuIdRef = ref(1);
  const menuScrollRef = ref(1);
  const dishesRef = ref([]);
  const tabRef = ref();
  function scroll({ detail: { scrollTop } }) {
    let offsetWindow = uni.getWindowInfo();

    const query = uni.createSelectorQuery().in(this);
    dishesRef.value.forEach(item => {
      query
        .select('#tab-' + item.id)
        .boundingClientRect(({ top, height }) => {
          if (
            top - offsetWindow.windowTop - 56 - 66.69 < 0 &&
            Math.abs(top - offsetWindow.windowTop - 56 - 66.69) < height
          ) {
            menuScrollRef.value = item.id;
            return;
          }
        })
        .exec();
    });
  }
  onLoad(() => {
    api.categort.getCategorys(true).then(res => {
      dishesRef.value = res.data;
    });
  });
</script>

<style lang="scss" scoped>
  @import '@/uni.scss';
  .container {
    display: flex;
    overflow: hidden;
  }
  .menu {
    width: 28%;
    padding: 0px 10px 0px 0px;
    display: flex;
    justify-items: center;
    flex-direction: column;
    align-items: center;
    background-color: rgb(246, 246, 246);
    .active {
      background-color: #ffffff;
      border-left: 8px solid $theme-color;
    }
    .item {
      width: 100%;
      height: 50px;
      display: flex;
      align-items: center;
      justify-content: left;
      .menuImg {
        width: 25px;
      }
    }
  }
  .menu,
  .sub-menu {
    height: 82vh;
    /* #ifdef H5 */
    height: calc(82vh - var(--window-top) - var(--window-bottom));
    /* #endif */
  }
  .sub-menu {
    background-color: rgb(255, 255, 255);
    width: 74%;
    .dishesItem {
      height: 80px;
    }
  }
  .menuItem {
    .title {
      height: 40px;
      display: flex;
      align-items: center;
      justify-content: left;
      .menuImg {
        width: 25px;
      }
    }
    .dishesItem {
      height: 120px;
      display: flex;
      align-items: center;
      .dishesImage {
        width: 60%;
        border-radius: 10px;
        margin: 8px;
      }
      .info {
        width: 80%;
        height: 120px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
      }
      .name {
        font-size: 1.2rem;
        font-weight: 600;
      }
      .jiage {
        height: 30px;
        font-size: 1rem;
        font-weight: 550;
        .jiage-qi {
          font-size: 0.7rem;
          font-weight: 500;
        }
      }
      .bo {
        display: flex;
        justify-content: space-between;
      }
      .jiagou {
        height: 25px;
        width: 48px;
        font-size: 15px;
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: $theme-color;
        color: #ffffff;
        padding: 0;
      }
    }
  }
</style>
