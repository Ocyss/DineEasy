<template>
  <uni-popup
    class="cart"
    ref="popupRef"
    type="bottom">
    <view class="content">666</view>
  </uni-popup>
  <view class="main">
    <view class="balance">
      <view @click="openCart">
        <uni-badge
          :text="quantityRef"
          absolute="rightTop"
          :customStyle="{ background: '#fff', color: '#000' }"
          size="normal">
          <uni-icons
            custom-prefix="custom-icon"
            type="cart"
            size="45"
            color="#fff" />
        </uni-badge>
      </view>

      ￥{{ priceRef }}
    </view>
    <view class="pay">去结算</view>
  </view>
</template>

<script setup>
  import { ref } from 'vue';
  const quantityRef = ref(0);
  const priceRef = ref(0);
  const cartRef = ref([]);
  const popupRef = ref();
  function openCart() {
    console.log(popupRef.value);
    popupRef.value.open('buttom');
    console.log(popupRef.value);
  }

  uni.$on('addPurchase', function (data) {
    console.log(data.quantity, data.attach, data.price);
    quantityRef.value += data.quantity;
    priceRef.value += data.price * data.quantity;
    priceRef.value += data.attach ? data.attach * data.quantity : 0;
    cartRef.value.push(data);
  });
</script>

<style scoped lang="scss">
  @import '@/uni.scss';
  .cart {
    .content {
      height: 70vh;
      background-color: #fff;
    }
    z-index: 9;
  }
  .main {
    width: 93%;
    position: fixed;
    bottom: calc(var(--window-bottom) + 5px);
    margin: 0 auto;
    left: 0;
    right: 0;
    color: #fff;
    display: flex;
    font-weight: 700;
    z-index: 999;
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
