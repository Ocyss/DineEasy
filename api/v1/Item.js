import Request from "@/api/request.js"
let request = new Request().http


/***
 * @description: 获取单品
 * @param . pagenum:偏移量, pagesize:数量,cid:分类id
 */
export function getItems(pagenum, pagesize, cid) {
  return request({
    url: "/api/v1/items",
    method: "GET",
    data: {
      pagenum,
      pagesize,
      cid
    }
  })
}

/***
 * @description: 增加单品
 * @param . name:单品名, description:单品描述, unit:单位, category_id:分类id, picture:图片, price:价格, status:状态
 */
export function addItem(data) {
  return request({
    url: "/api/v1/item/add",
    method: "POST",
    data
  })
}
