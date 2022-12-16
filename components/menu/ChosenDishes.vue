<template>
  <view class="main">
    <view class="balance">
      <uni-badge
        :text="1"
        absolute="rightTop"
        :customStyle="{ background: '#fff', color: '#000' }"
        size="normal">
        <uni-icons
          custom-prefix="custom-icon"
          type="cart"
          size="45"
          color="#fff" />
      </uni-badge>

      ￥11.33
    </view>
    <view class="pay">去结算</view>
  </view>
</template>

<script setup>
  import { ref } from 'vue';
</script>

<style scoped lang="scss">
  @import '@/uni.scss';

  .main {
    width: 93%;
    position: fixed;
    bottom: 12px;
    margin: 0 auto;
    left: 0;
    right: 0;
    color: #fff;
    display: flex;
    font-weight: 700;
  }
  .balance {
    height: 10vh;
    background-color: #000;
    border-radius: 15px 0 0 15px;
    width: 70%;
    display: flex;
    align-items: center;
    justify-content: space-evenly;
    font-size: 1.2rem;
  }
  .pay {
    width: 30%;
    background-color: $theme-color;
    border-radius: 0 15px 15px 0;
    height: 10vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 1.4rem;
  }
</style>
