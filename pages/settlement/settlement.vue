<template>
  <view class="main">
    <view class="content">
      <view
        v-for="itemtype in itemData.cart"
        :key="itemtype">
        <view
          class="con"
          v-for="item in itemtype"
          :key="item">
          <view class="left">
            <image
              class="image"
              :src="item.picture"
              mode="aspectFill"></image>
            <view class="info">
              <view class="name">{{ item.name }}</view>
              <view class="quantity">x{{ item.quantity }}</view>
            </view>
          </view>
          <view class="right">￥{{ item.price }}</view>
        </view>
      </view>
    </view>
  </view>
  <view class="submit">
    <view class="cnt">合计: ￥{{ itemData.price }}</view>
    <view class="button">
      <button
        type="primary"
        @click="toPay">
        付款
      </button>
    </view>
  </view>
</template>

<script setup>
  import { ref } from 'vue';
  const itemData = ref({});
  function toPay() {
    uni.navigateTo({ url: '/pages/pay/success' });
  }
  uni.$on('startSettlement', function (data) {
    itemData.value = data;
  });
</script>

<style scoped lang="scss">
  .main {
    padding: 15px;
  }
  .content {
    margin-bottom: 100px;
  }
  .con {
    margin: 12px 0;
    display: flex;
    align-items: center;
    justify-content: space-between;
    .left {
      display: flex;
      .image {
        width: 70px;
        height: 60px;
        border-radius: 5px;
      }
      .info {
        margin-left: 8px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        .name {
          font-weight: 600;
          font-size: 16px;
        }
        .quantity {
          color: $theme-color;
        }
      }
    }
    .right {
      font-size: 18px;
      font-weight: 600;
    }
  }
  .submit {
    position: fixed;
    bottom: 0px;
    width: 100%;
    height: 10vh;
    color: $theme-color;
    background-color: #fff;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 25px;
    font-weight: 600;
    font-size: 19px;
    box-sizing: border-box;
    filter: drop-shadow(0 5px 10px #000);
    .button {
      width: 30%;
    }
  }
</style>
