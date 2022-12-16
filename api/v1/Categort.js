import Request from "@/api/request.js"
let request = new Request().http


/***
 * @description: 获取分类
 * @param bool load 是否加载单品信息
 */
export function getCategorys(load) {
  return request({
    url: "/api/v1/categorys",
    method: "GET",
    data: {
      load
    }
  })
}

/***
 * @description: 增加分类
 * @param . name:分类名, picture:图片, status:分类状态
 */
export function addCategory(data) {
  return request({
    url: "/api/v1/category/add",
    method: "POST",
    data
  })
}
