let baseUrl = "http://192.168.0.109:3001"

export default class Request {
  http(param) {
    // 请求参数
    var url = param.url,
      method = param.method,
      header = {
        ...param.header,
        'content-type': "application/json"
      },
      data = param.data || {},
      token = param.token || "",
      hideLoading = param.hideLoading || false;

    //拼接完整请求地址
    var requestUrl = baseUrl + url;

    //加载圈
    if (!hideLoading) {
      uni.showLoading({
        title: '加载中...'
      });
    }

    // 返回promise
    return new Promise((resolve, reject) => {
      // 请求
      uni.request({
        url: requestUrl,
        data: data,
        method: method,
        header: header,
        success: (res) => {
          // 判断 请求api 格式是否正确
          if (res.statusCode && res.statusCode != 200 && res.statusCode != 400) {
            uni.showToast({
              title: "api错误" + res.errMsg,
              icon: 'none'
            });
            return;
          }
          // code判断:200成功,不等于200错误
          if (res.data.code) {
            if (res.data.code != '200') {
              uni.showToast({
                title: "错误:" + res.data.msg,
                icon: 'none'
              });
              return;
            }
          } else {
            uni.showToast({
              title: "返回格式错误！",
              icon: 'none'
            });
            return;
          }
          // 将结果抛出
          resolve(res.data)
        },
        //请求失败
        fail: (e) => {
          console.log(e)
          uni.showToast({
            title: "" + e.errMsg,
            icon: 'none'
          });
          resolve({
            code: 444,
            msg: "请求失败!"
          });
        },
        //请求完成
        complete() {
          //隐藏加载
          if (!hideLoading) {
            uni.hideLoading();
          }
          resolve();
          return;
        }
      })
    })
  }
}
