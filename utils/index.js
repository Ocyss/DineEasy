//获取当天的分钟数
export function getTimeMinutes() {
  let date = new Date()
  return date.getMinutes() + date.getHours() * 60
}

//分钟数转时间
export function hToMinute(h) {
  let hstr = (Math.floor(h / 60)).toString()
  let mstr = (h % 60).toString()
  return (hstr.length < 2 ? "0" + hstr :
    hstr) + ":" + (mstr.length < 2 ? "0" + mstr : mstr);
}

export function isItOpen(sTime, eTime) {
  let minutes = getTimeMinutes()
  return (sTime < minutes && minutes < eTime)
}
