<template>
  <uni-popup
    class="popup"
    :safe-area="false"
    @maskClick="closeNorm"
    ref="popupRef">
    <scroll-view
      class="popupContent"
      scroll-y>
      <image
        class="popupImage"
        :src="cGroup.picture"
        mode="aspectFill" />
      <view class="popupText">
        <view class="title">{{ cGroup.name }}</view>
        <view class="description">{{ cGroup.description }}</view>
      </view>

      <view
        v-for="groupItem in cGroup.combo_group"
        :key="groupItem.id">
        <uni-section
          :title="groupItem.name"
          type="line">
          <TagList
            style="margin: 2px 10px"
            :limit="groupItem.quantity"
            :tagList="groupItem.items"
            :fixed="groupItem.fixed"
            :same="groupItem.same"
            :selectedTagList="seletedTagListRef['group'][groupItem.id]"
            @select="d => handleSelectTag(d, groupItem)" />
        </uni-section>
      </view>
    </scroll-view>
    <view class="popupCover">
      <view class="info">
        <view>
          <view class="price">
            ￥{{ (cGroup.price + seletedTagListRef['attach']) * seletedTagListRef.quantity }}
          </view>
          <view class="content">
            <text
              v-for="group in seletedTagListRef['group']"
              :key="group">
              <text
                v-for="item in group"
                :key="item.id">
                {{ item.name }}
                <text v-if="item.quantity != 1">*{{ item.quantity }}</text>
                &nbsp;
              </text>
            </text>
          </view>
        </view>
        <view>
          <uni-number-box
            v-model="seletedTagListRef.quantity"
            :min="1"
            :max="99" />
        </view>
      </view>
      <view class="button">
        <button
          plain
          class="qx"
          @click="closeNorm">
          取消
        </button>
        <button
          plain
          class="qr"
          @click="sure">
          确认
        </button>
      </view>
    </view>
  </uni-popup>
</template>

<script setup>
  import { ref, toRefs } from 'vue';
  import TagList from '@/components/tag-list/tag-list.vue';
  // const props = defineProps(['cGroup']);
  // const { cGroup } = toRefs(props);
  const cGroup = ref({});
  const popupRef = ref();
  const seletedTagListRef = ref({
    id: 0,
    price: 0,
    group: {},
    attach: 0,
    quantity: 1
  });
  const selectNorm = item => {
    cGroup.value = item;
    seletedTagListRef.value = {
      id: cGroup.value.id,
      price: cGroup.value.price,
      group: {},
      attach: 0,
      quantity: 1
    };
    uni.hideTabBar({ animation: true });
    popupRef.value.open('bottom');
  };

  const closeNorm = () => {
    popupRef.value.close();
    setTimeout(() => {
      uni.showTabBar({ animation: true });
    }, 100);
  };

  const handleSelectTag = (seletedTag, groupItem) => {
    seletedTagListRef.value['group'][groupItem.id] = seletedTag;
    seletedTagListRef.value['attach'] = Object.values(seletedTagListRef.value['group']).reduce(
      (total, group) => {
        return total + group.reduce((subtotal, g) => subtotal + g.price, 0);
      },
      0
    );
  };

  const sure = () => {
    if (
      cGroup.value['combo_group'].every(item => {
        if (!item.fixed && seletedTagListRef.value['group'][item.id].length !== item.quantity) {
          uni.showToast({
            title: `${item.name}还没选呢!`,
            icon: 'none'
          });
          return false;
        }
        return true;
      })
    ) {
      uni.$emit('addPurchase', seletedTagListRef.value);
      closeNorm();
    }
  };
  defineExpose({
    selectNorm
  });
</script>

<style scoped lang="scss">
  @import '@/uni.scss';
  .popupContent {
    background-color: #fff;
    width: 100vw;
    height: 60vh;
    border-radius: 25px 25px 0 0;
    overflow: hidden;
    .popupImage {
      width: 100%;
      height: 250px;
    }
    .popupText {
      padding: 15px;
      .title {
        font-size: 23px;
        font-weight: 700;
      }
      .description {
        font-size: 13px;
        color: #5f5f5f;
      }
    }
  }
  .popupCover {
    width: 100vw;
    min-height: 120px;
    background-color: #ffffff;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    padding: 5px 20px;
    box-sizing: border-box;
    box-shadow: 0px -20px 10px rgba(0, 0, 0, 0.9);
    .button {
      display: flex;
      align-items: center;
      uni-button {
        font-size: unset;
        display: flex;
        align-items: center;
        justify-content: center;
        &.qr {
          background-color: $theme-color;
          color: #fff;
        }
        height: 92%;
        border-color: $theme-color;
        width: 48%;
        color: $theme-color;
        border-radius: 25px;
      }
    }
    .info {
      display: flex;
      align-items: center;
      justify-content: space-between;
      .uni-numbox {
        color: $theme-color;
        :deep(.uni-numbox__minus) {
          border-radius: 25px 0 0 25px;
        }
        :deep(.uni-numbox__plus) {
          border-radius: 0 25px 25px 0;
        }
        :deep(.uni-numbox__value),
        :deep(.uni-numbox__plus),
        :deep(.uni-numbox__minus) {
          margin: 0 !important;
          color: #000 !important;
          font-size: 18px;
          border: 1px solid #616161;
          box-sizing: border-box;
          height: 32px;
        }
        :deep(.uni-numbox__value) {
          border-left-style: none;
          border-right-style: none;
        }
        :deep(.uni-numbox-btns) {
          uni-text {
            color: $theme-color !important;
          }
        }
      }
      .content {
        font-size: 12px;
        color: #181818;
      }
      .price {
        font-size: 18px;
        font-weight: 600;
      }
    }
  }
</style>
