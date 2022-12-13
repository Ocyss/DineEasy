<template>
  <view class="container">
    <scroll-view scroll-y scroll-with-animation class="menu">
      <view
        class="item"
        :class="{ active: menuIdRef === item.id }"
        v-for="item in menuRef"
        :key="item"
        @click="menuIdRef = item.id"
      >
        {{ item.name }}
      </view>
    </scroll-view>
    <scroll-view
      class="sub-menu"
      scroll-y
      :scroll-into-view="'tab-' + menuIdRef"
      scroll-with-animation
      @scroll="handleScroll"
      show-scrollbar="true"
    >
      <view
        class="menuItem"
        v-for="menuItem in dishesRef"
        :id="'tab-' + menuItem.menuId"
        :key="menuItem.menuId"
      >
        <view class="title">******{{ menuItem.menuName }}******</view>
        <view
          class="dishesItem"
          v-for="dishesItem in menuItem.data"
          :key="dishesItem"
        >
          <image
            class="dishesImage"
            src="../../static/333.png"
            mode="widthFix"
          ></image>
          <view class="info">
            <view class="name">{{ dishesItem }}</view>
            <view class="bo">
              <view class="jiage">
                ￥11.11
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
const menuIdRef = ref(0);
const menuRef = ref([
  { id: 0, name: '包点' },
  { id: 1, name: '糯米饭' },
  { id: 2, name: '面点' },
  { id: 3, name: '豪华套餐' },
  { id: 4, name: '额外小料' }
]);
const dishesRef = ref([
  {
    menuId: 0,
    menuName: '包点',
    data: ['肉包', '叉烧包', '粉丝包', '馒头', '强哥包']
  },
  {
    menuId: 1,
    menuName: '糯米饭',
    data: ['咸糯米饭', '甜糯米饭', '香糯米饭']
  },
  { menuId: 2, menuName: '面点', data: ['圆面', '细面', '下面', '牛肉面'] },
  {
    menuId: 3,
    menuName: '豪华套餐',
    data: ['肉包套餐', '叉烧包糯米饭套餐', '强哥蛋白质套餐']
  },
  {
    menuId: 4,
    menuName: '额外小料',
    data: ['蛋白质', '强哥', '菜籽', '面粉']
  }
]);

function handleScroll() {
  let offsetWindow = uni.getWindowInfo();

  const query = uni.createSelectorQuery().in(this);
  menuRef.value.forEach(item => {
    query
      .select('#tab-' + item.id)
      .boundingClientRect(({ top, height }) => {
        if (
          top - offsetWindow.windowTop - 56 - 66.69 < 0 &&
          Math.abs(top - offsetWindow.windowTop - 56 - 66.69) < height
        ) {
          menuIdRef.value = item.id;
        }
      })
      .exec();
  });
}
</script>

<style lang="scss" scoped>
@import '@/uni.scss';
.container {
  display: flex;
  overflow: hidden;
}
.menu {
  width: 26%;
  display: flex;
  justify-items: center;
  flex-direction: column;
  align-items: center;
  view {
    height: 30px;
  }
  background-color: rgb(246, 246, 246);
  .active {
    background-color: #ffffff;
    border-left: 2px solid $theme-color;
  }
  .item {
    height: 50px;
    display: flex;
    align-items: center;
    justify-content: center;
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
  }
  .dishesItem {
    height: 120px;
    display: flex;
    align-items: center;
    .dishesImage {
      width: 60%;
    }
    .info {
      width: 80%;
      height: 120px;
      display: flex;
      flex-direction: column;
      justify-content: space-around;
    }
    .name {
      font-size: 1.2rem;
      font-weight: 600;
    }
    .jiage {
      font-size: 1rem;
      font-weight: 550;
      .jiage-qi {
        font-size: 0.6rem;
        font-weight: 500;
      }
    }
    .bo {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }
    .jiagou {
      height: 30px;
      display: flex;
      align-items: center;
      background-color: $theme-color;
      color: #ffffff;
    }
  }
}
</style>
