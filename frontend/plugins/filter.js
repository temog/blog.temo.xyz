import Vue from 'vue'

Vue.filter('date', (value) => {
  const date = new Date(value)
  const day = ['SUN', 'MON', 'TUE', 'WED', 'THU', 'FRI', 'SAT']

  const y = date.getFullYear()
  const m = ('00' + (date.getMonth() + 1)).slice(-2)
  const d = ('00' + date.getDate()).slice(-2)
  const dow = day[date.getDay()]
  const h = ('00' + date.getHours()).slice(-2)
  const i = ('00' + date.getMinutes()).slice(-2)
  return `${y}.${m}.${d} ${dow} ${h}:${i}`
})
