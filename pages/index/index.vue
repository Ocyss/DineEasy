<template>
    <view class="main">
        <image src="~@/static/111.jpeg" alt="" class="bgimg" />
        <view class="title">三姐弟</view>
        <view class="content">
            <view class="wdhy">
                <image src="~@/static/index/hy.png" class="hyimg" />
                <view class="">我的会员</view>
            </view>
            <view class="wdqb">
                <image src="~@/static/index/qb.png" class="qbimg" />
                <view class="">我的券包</view>
            </view>
        </view>
        <view class="members">
            <view class="hysc">
                <image src="~@/static/index/sc.png" class="scimg" />
                会员商城
            </view>
            <view class="meri">
                <view class="hhr">
                    合伙人
                    <image src="~@/static/index/hhr.png" class="hhrimg" />
                </view>
                <view class="cz">
                    特惠充值
                    <image src="~@/static/index/cz.png" class="czimg" />
                </view>
            </view>
        </view>
        <image src="~@/static/222.jpeg" alt="" class="bgimg" />
    </view>
</template>

<script setup></script>

<style lang="scss" scoped>
.main {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
}

.bgimg {
    width: 100%;
}

.content {
    width: 90%;
    box-sizing: border-box;
    background-color: #fff;
    border: 2px solid #e2e2e2;
    padding: 1rem;
    display: flex;
    align-items: center;
    justify-content: space-around;
    font-size: 1.3rem;
    font-weight: 800;
    position: relative;
    top: -1.5rem;
    border-radius: 0.8rem;
    .wdhy,
    .wdqb {
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
    }
    .qbimg,
    .hyimg {
        width: 50px;
        height: 50px;
    }
}
.members {
    width: 90%;
    box-sizing: border-box;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 1.2rem;
    font-weight: 700;
    margin-bottom: 2rem;
    .scimg {
        width: 50px;
        height: 50px;
    }
    .hhrimg,
    .czimg {
        width: 43px;
        height: 43px;
    }
    .hysc {
        display: flex;
        align-items: center;
        justify-content: space-evenly;
        border: 2px solid #e2e2e2;
        height: 110px;
        width: 48%;
        border-radius: 0.8rem;
    }
    .meri {
        width: 48%;
        display: flex;
        height: 110px;
        flex-direction: column;
        justify-content: space-between;
    }
    .hhr,
    .cz {
        font-weight: 500;
        font-size: 1rem;
        box-sizing: border-box;
        display: flex;
        align-items: center;
        justify-content: space-between;
        border: 2px solid #e2e2e2;
        width: 100%;
        height: 50px;
        padding: 0.3rem;
        border-radius: 0.8rem;
    }
}
.title {
    font-size: 4rem;
    font-weight: 800;
    color: #fff;
    position: absolute;
    top: 4rem;
}
</style>
