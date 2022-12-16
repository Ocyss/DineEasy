import Request from "@/api/request.js"
let request = new Request().http


/***
 * @description: 上传图片
 * @param . file:文件,class:类别(category,item)
 */
export function upLoad(file, fileClass) {
  return request({
    url: "/api/v1/upload",
    method: "POST",
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data: {
      file,
      class: fileClass
    }
  })
}
