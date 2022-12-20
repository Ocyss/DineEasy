<template>
  <view class="tag-list">
    <view
      class="tag-item"
      :class="getStyle(item.id)"
      v-for="(item, index) in tagList"
      :key="item.id"
      @click="handleSelectTag(item, index)">
      {{ item['name'] }}*{{ item['quantity'] }}
      <text
        class="price"
        v-if="item['price'] != 0">
        ￥{{ item['price'] }}
      </text>
    </view>
  </view>
</template>

<script setup>
  import { toRefs, defineProps, ref, defineEmits } from 'vue';

  const emit = defineEmits(['select']);
  const props = defineProps({
    //选中tag列表
    selectedTagList: {
      type: Array,
      default: () => []
    },
    //展示tag列表
    tagList: {
      type: Array,
      default: () => []
    },
    //限制tag选中个数
    limit: {
      type: Number,
      default: 1
    },
    //是否固定
    fixed: {
      type: Boolean,
      default: false
    },
    //是否可选同种单品
    same: {
      type: Boolean,
      default: false
    }
  });
  const { selectedTagList, tagList, limit, fixed, same } = toRefs(props);
  if (fixed.value) {
    emit('select', tagList.value);
  }
  /**
   * 获取tag的样式
   */
  function getStyle(id) {
    return {
      seleted: fixed.value || getIndexSeleted(id) !== -1,
      fixed: fixed.value
    };
  }
  /**
   * @param {Object} item 点击tag的每一项
   * @param {Object} index 索引
   */
  function handleSelectTag(item, index) {
    if (fixed.value) {
      return;
    }
    let sList = JSON.parse(JSON.stringify(selectedTagList.value));
    const indexSelected = this.getIndexSeleted(item.id);
    if (indexSelected !== -1) {
      sList.splice(indexSelected, 1);
    } else if (this.limit == 1) {
      sList = [item];
    } else if (selectedTagList.value.length == this.limit) {
      uni.showToast({
        title: `最多能选择${this.limit}个哦!`,
        icon: 'none'
      });
    } else {
      sList.push(item);
    }
    emit('select', sList);
  }
  /**
   * 获取tag的选中索引
   */
  function getIndexSeleted(id) {
    return selectedTagList.value.findIndex(item => item.id == id);
  }
</script>

<style lang="scss" scoped>
  @import '@/uni.scss';
  .tag-list {
    display: flex;
    flex-wrap: wrap;

    .tag-item {
      color: #333;
      font-size: 26rpx;
      background: rgba(230, 230, 230, 1);
      padding: 8rpx 28rpx;
      border-radius: 8rpx;
      margin: 0 20rpx 20rpx 0;
      position: relative;
      .price {
        font-weight: 600;
        color: $theme-color;
      }
      &.seleted {
        background: linear-gradient(to right, $theme-color, $theme-color);
        color: #fff;
        .price {
          color: #fff;
        }
      }
    }
  }
</style>
