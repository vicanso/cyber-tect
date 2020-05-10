import { sha256 } from '@/helpers/crypto'

const hash = 'JT'

export function generatePassword (pass) {
  return sha256(pass + hash)
}

export function formatDate (str) {
  const d = new Date(str)
  return `${d.toLocaleDateString()} ${d.toLocaleTimeString()}`
}

export function formatDuration (d) {
  if (d > 1000) {
    const v = d / 100
    let fix = 1
    if (Number.parseInt(v) % 10 === 0) {
      fix = 0
    }
    return `${(d / 1000).toFixed(fix)} 秒`
  }
  return `${d} 毫秒`
}

export function delay (ms) {
  return new Promise(resolve => setTimeout(resolve, ms))
}

export function diff (current, original) {
  const data = {}
  let modifiedCount = 0
  Object.keys(current).forEach((key) => {
    if (current[key] !== original[key]) {
      data[key] = current[key]
      modifiedCount++
    }
  })
  return {
    modifiedCount,
    data
  }
}
