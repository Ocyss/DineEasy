import Request from "@/api/request.js"
let request = new Request().http


/***
 * @description: 获取门店
 * @param . pagenum:偏移量, pagesize:数量
 */
export function getStores(pagenum, pagesize) {
  return request({
    url: "/api/v1/stores",
    method: "GET",
    data: {
      pagenum,
      pagesize
    }
  })
}


/***
 * @description: 增加门店
 * @param . name:门店名字,address:门店地址,phone:门店电话,picture:门店照片,startTime:营业开始时间,endTime:营业结束时间,status:门店状态
 */
export function addStore(data) {
  return request({
    url: "/api/v1/stores",
    method: "POST",
    data
  })
}
