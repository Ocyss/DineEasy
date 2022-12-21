<template>
  <uni-popup
    class="cart"
    ref="popupRef"
    type="bottom"
    @change="popupChange">
    <view class="popupTop">
      <view>全选</view>
      <view>清空</view>
    </view>
    <scroll-view
      class="popupContent"
      scroll-y>
      <uni-section
        title="套餐"
        type="line">
        <view
          v-for="combo in cartRef.combo"
          :key="combo">
          <view class="com">
            <view class="left">
              <radio
                style="transform: scale(0.7)"
                color="#dc5711"
                :checked="combo.checked"
                @click="combo.checked = !combo.checked" />
              <image
                class="image"
                :src="combo.picture"
                mode="aspectFill" />
              <view class="info">
                <view class="name">{{ combo.name }}</view>
                <view class="price">￥{{ combo.price }}</view>
              </view>
            </view>
            <view class="right">
              <uni-number-box
                style="transform: scale(0.9)"
                :value="combo.quantity"
                background="#dc5711"
                color="#fff" />
            </view>
          </view>
        </view>
      </uni-section>
      <uni-section
        title="单品"
        type="line">
        <view
          class=""
          v-for="item in cartRef.item"
          :key="item">
          <view class="com">
            <view class="left">
              <radio
                style="transform: scale(0.7)"
                color="#dc5711"
                :checked="item.checked"
                @click="item.checked = !item.checked" />
              <image
                class="image"
                :src="item.picture"
                mode="aspectFill" />
              <view class="info">
                <view class="name">{{ item.name }}</view>
                <view class="price">￥{{ item.price }}</view>
              </view>
            </view>
            <view class="right">
              <uni-number-box
                style="transform: scale(0.9)"
                :value="item.quantity"
                background="#dc5711"
                color="#fff" />
            </view>
          </view>
        </view>
      </uni-section>
    </scroll-view>
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
    <view
      class="pay"
      @click="toSettlement">
      去结算
    </view>
  </view>
</template>

<script setup>
  import { ref } from 'vue';
  const quantityRef = ref(0);
  const priceRef = ref(0);
  const cartRef = ref({ item: {}, combo: {} });
  const popupRef = ref();
  let showPopup = false;
  function openCart() {
    if (!showPopup) {
      popupRef.value.open('buttom');
    } else {
      popupRef.value.close();
    }
  }

  function popupChange({ show, type }) {
    showPopup = show;
  }
  function toSettlement() {
    if (quantityRef.value > 0) {
      uni.navigateTo({ url: '/pages/settlement/settlement' });
      setTimeout(() => {
        uni.$emit('startSettlement', {
          quantity: quantityRef.value,
          price: priceRef.value,
          cart: cartRef.value
        });
      }, 300);
    } else {
      uni.showToast({
        title: `请先选择单品哦!`,
        icon: 'none'
      });
    }
  }
  uni.$on('addPurchase', function ({ type, data }) {
    let price = (data.price + data.attach) * data.quantity;

    quantityRef.value += data.quantity;
    priceRef.value += price;
    if (cartRef.value[type][data.id]) {
      cartRef.value[type][data.id].quantity += data.quantity;
      cartRef.value[type][data.id].price += price;
    } else {
      cartRef.value[type][data.id] = {
        name: data.name,
        unit: data.unit,
        quantity: data.quantity,
        picture: data.picture,
        price: price,
        checked: true
      };
    }
  });
</script>

<style scoped lang="scss">
  @import '@/uni.scss';
  .com {
    display: flex;
    margin-top: 2px;
    height: 60px;
    margin-bottom: 13px;
    .left {
      display: flex;
      align-items: center;
      width: 70%;
      .switch {
      }
      .image {
        width: 80px;
        height: 60px;
        border-radius: 8px;
      }
      .info {
        display: flex;
        height: 60px;
        flex-direction: column;
        justify-content: space-between;
      }
    }
    .right {
      width: 30%;
      display: flex;
      align-items: center;
    }
  }
  .cart {
    .popupTop {
      border-radius: 25px 25px 0 0;
      overflow: hidden;
      display: flex;
      background-color: #dedede;
      height: 5vh;
      align-items: center;
      justify-content: space-between;
      padding: 20px;
      box-sizing: border-box;
    }
    .popupContent {
      background-color: #fff;
      height: 60vh;
      padding: 15px;
      padding-top: 0;
      box-sizing: border-box;
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
    z-index: 60;
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
