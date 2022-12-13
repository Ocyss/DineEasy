<template>
  <view>
    <view class="search-input">
      <uni-search-bar
        @input="inputend"
        cancelButton="none"
        clearButton="none"
        placeholder="请输入店名或者地址"
        :radius="100" />
    </view>
    <view class="store-list">
      <view
        class="store-item"
        @click="selectStore(item)"
        :class="selectId == item.id ? 'select' : ''"
        v-for="item in storeDataRef"
        :key="item">
        <view class="store-item-info">
          <view class="name">{{ item.name }}</view>
          <view class="type">
            <view v-if="item.state.proofing">
              <uni-tag
                text="打样中"
                type="error"
                inverted="true"
                size="mini" />
            </view>
            <view v-if="item.state.dineIn">
              <uni-tag
                text="可堂食"
                type="success"
                inverted="true"
                size="mini" />
            </view>
            <view v-if="item.state.pickUp">
              <uni-tag
                text="可自提"
                type="primary"
                inverted="true"
                size="mini" />
            </view>
            <view v-if="item.state.takeaway">
              <uni-tag
                text="可外卖"
                type="warning"
                inverted="true"
                size="mini" />
            </view>
          </view>
          <view class="address">
            <uni-icons
              custom-prefix="iconfont"
              type="icon-wxbdingwei"
              size="20" />
            {{ item.address }}
          </view>
          <view class="startTime">
            <uni-icons
              custom-prefix="iconfont"
              type="icon-shijian"
              size="20" />
            {{ item.startTime }}-{{ item.endTime }}
          </view>
        </view>
        <view class="store-item-operate">
          <view>去点单</view>
          <view class="jl">距离{{ item.distance }}米</view>
          <view class="operate">
            <view>
              <uni-icons
                color="#f0ad4e"
                custom-prefix="iconfont"
                type="icon-dianhuazhengzaibohao"
                size="23" />
            </view>
            <view>
              <uni-icons
                color="#f0ad4e"
                custom-prefix="iconfont"
                type="icon-daohang"
                size="23" />
            </view>
          </view>
        </view>
      </view>
    </view>
  </view>
</template>

<script setup>
  import { ref } from 'vue';
  import { useConfig } from '@/store/config.js';

  const config = useConfig();
  const selectId = ref(config.Store.id);
  //状态 0打样，1可堂食，2可自提，3可外卖
  const storeDataRef = ref([
    {
      id: 1,
      name: '三姐弟大南店',
      state: {
        proofing: false,
        dineIn: true,
        pickUp: true,
        takeaway: false
      },
      startTime: '05:30',
      endTime: '13:30',
      address: '浙江省温州市鹿城区大南街146号',
      distance: '1234'
    },
    {
      id: 2,
      name: '三姐弟大呃店',
      state: {
        proofing: false,
        dineIn: true,
        pickUp: true,
        takeaway: false
      },
      startTime: '05:30',
      endTime: '13:30',
      address: '浙江省温州市鹿城区梨泰街146号',
      distance: '645'
    },
    {
      id: 3,
      name: '三姐弟大被店',
      state: {
        proofing: false,
        dineIn: true,
        pickUp: true,
        takeaway: false
      },
      startTime: '05:30',
      endTime: '13:30',
      address: '浙江省温州市鹿城区复打街146号',
      distance: '789'
    },
    {
      id: 4,
      name: '三姐弟大都店',
      state: {
        proofing: false,
        dineIn: true,
        pickUp: true,
        takeaway: false
      },
      startTime: '05:30',
      endTime: '13:30',
      address: '浙江省温州市鹿城区回放街146号',
      distance: '45645'
    },
    {
      id: 5,
      name: '三姐弟大卡店',
      state: {
        proofing: false,
        dineIn: true,
        pickUp: true,
        takeaway: false
      },
      startTime: '05:30',
      endTime: '13:30',
      address: '浙江省温州市鹿城区便宜街146号',
      distance: '43123'
    },
    {
      id: 6,
      name: '三姐弟雪上店',
      state: {
        proofing: false,
        dineIn: true,
        pickUp: true,
        takeaway: false
      },
      startTime: '05:30',
      endTime: '13:30',
      address: '浙江省温州市鹿城区集火街146号',
      distance: '12345'
    },
    {
      id: 7,
      name: '三姐弟萨达店',
      state: {
        proofing: false,
        dineIn: true,
        pickUp: true,
        takeaway: false
      },
      startTime: '05:30',
      endTime: '13:30',
      address: '浙江省温州市鹿城区大师街146号',
      distance: '456'
    }
  ]);
  let inputTime = null;

  //搜索门店
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

  function selectStore(item) {
    if (item.id == selectId.value) {
      uni.switchTab({
        url: '/pages/order/order'
      });
    }
    selectId.value = item.id;
    config.Store.id = item.id;
    config.Store.name = item.name;
    config.Store.address = item.address;
  }
</script>

<style lang="scss">
  @import '@/uni.scss';

  .store-list {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }

  .store-item {
    height: 10rem;
    border: 2px solid #e2e2e2;
    width: 90%;
    padding: 1.1rem;
    padding-right: 0px;
    box-sizing: border-box;
    margin-bottom: 1rem;
    border-radius: 0.7rem;
    display: flex;
    justify-content: space-between;
  }

  .store-item-info {
    .name {
      font-size: 1.2rem;
      font-weight: 600;
    }

    .type {
      display: flex;

      view {
        margin: 3px;
      }
    }

    .address,
    .startTime {
      font-size: 0.7rem;
      display: flex;
      align-items: center;
    }
  }

  .store-item-operate {
    width: 38%;
    display: flex;
    align-items: center;
    flex-direction: column;
    border-left-style: solid;
    border-left-width: 1px;
    border-left-color: #c3c3c3;
    font-size: 0.9rem;
    padding: 1.1rem;
    box-sizing: border-box;
    justify-content: space-between;

    .jl {
      font-size: 0.7rem;
    }

    .operate {
      display: flex;
      width: 100%;

      justify-content: space-around;

      view {
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: rgb(234, 234, 234);
        width: 36px;
        height: 36px;
        border-radius: 50%;
      }
    }
  }

  .select {
    border: 3px solid $theme-color;
    position: relative;
    box-shadow: 0px 2px 7px 0px rgba(85, 110, 97, 0.35);

    &:before {
      content: '';
      position: absolute;
      right: 0;
      bottom: 0;
      border: 16px solid $theme-color;
      border-top-color: transparent;
      border-left-color: transparent;
    }

    &:after {
      content: '';
      width: 4px;
      height: 14px;
      position: absolute;
      right: 6px;
      bottom: 3px;
      border: 2px solid #fff;
      border-top-color: transparent;
      border-left-color: transparent;
      transform: rotate(45deg);
    }
  }
</style>
