import Request from "@/api/request.js"
let request = new Request().http


/***
 * @description: 获取套餐
 */
export function getCombos() {
  return request({
    url: "/api/v1/combos",
    method: "GET",
  })
}
