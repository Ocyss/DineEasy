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
      <view style="height: 200px" />
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
            lazy-load
            :src="dishesItem.picture"
            mode="aspectFill"></image>
          <view class="info">
            <view>
              <view class="name">{{ dishesItem.name }}</view>
              <view class="description">
                {{ dishesItem.description }}
              </view>
            </view>
            <view class="bo">
              <view class="jiage">
                <text class="jiage-fh">￥</text>
                <text>{{ dishesItem.price }}</text>
                <text class="jiage-qi">&nbsp;起</text>
              </view>
              <button
                v-if="dishesItem.combo_group"
                size="mini"
                class="jiagou"
                @click="selectNorm(dishesItem)">
                选规格
              </button>
              <uni-icons
                v-else
                class="jh"
                type="plus-filled"
                size="35"></uni-icons>
            </view>
          </view>
        </view>
      </view>
    </scroll-view>
    <uni-popup
      ref="popupRef"
      background-color="#fff"
      type="dialog">
      <view class="popupContent"></view>
    </uni-popup>
  </view>
</template>

<script setup>
  import { ref } from 'vue';
  import { onLoad } from '@dcloudio/uni-app';
  import api from '@/api';
  import { getCurrentInstance } from 'vue';

  const instance = getCurrentInstance();
  const menuIdRef = ref(1);
  const menuScrollRef = ref(1);
  const dishesRef = ref([]);
  const tabRef = ref();
  const popupRef = ref();
  function scroll({ detail: { scrollTop } }) {
    let offsetWindow = uni.getWindowInfo();
    let leftHeight = 50;
    const query = uni.createSelectorQuery().in(instance);
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
  function selectNorm(item) {
    popupRef.value.open('center');
  }

  onLoad(() => {
    api.categort.getCategorys(true).then(res => {
      dishesRef.value = res.data;
    });
    api.combo.getCombos().then(res => {
      dishesRef.value.forEach(item => {
        if (item.id == 1) {
          item.items.push(...res.data);
          return;
        }
      });
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
    width: 24%;
    margin: 0px 10px 0px 0px;
    display: flex;
    font-size: 0.9rem;
    justify-items: center;
    flex-direction: column;
    align-items: center;
    background-color: rgb(233, 233, 233);
    .active {
      background-color: #ffffff;
    }
    .active::after {
      content: '';
      background-color: $theme-color;
      width: 6px;
      height: 50px;
      left: 0;
      position: absolute;
    }
    .item {
      width: 100%;
      height: 50px;
      display: flex;
      justify-content: center;
      align-items: center;
      flex-direction: column;
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
      margin-bottom: 10px;
      .dishesImage {
        width: 62%;
        height: 120px;
        border-radius: 10px;
        margin: 8px;
      }
      .info {
        width: 80%;
        height: 120px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        padding: 0 5px 0 0;
      }
      .name {
        font-size: 1.1rem;
        font-weight: 600;
      }
      .description {
        font-size: 0.8rem;
        color: #5f5f5f;
      }
      .jiage {
        font-size: 1.1rem;
        font-weight: 550;
        display: inline;
        .jiage-qi {
          font-size: 0.7rem;
          font-weight: 500;
          color: #5f5f5f;
        }
        .jiage-fh {
          font-size: 0.7rem;
        }
      }
      .bo {
        display: flex;
        align-items: center;
        justify-content: space-between;
      }
      .jiagou {
        width: 50px;
        margin: 0;
        background-color: $theme-color;
        color: #ffffff;
        padding: 0;
        border-radius: 25px;
      }
      .jh {
        color: $theme-color !important;
        cursor: pointer;
      }
    }
  }
</style>
