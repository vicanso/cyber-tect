import { sha256 } from '@/helpers/crypto'

const hash = 'JT'

export function generatePassword (pass) {
  return sha256(pass + hash)
}

export function formatDate (str) {
  const d = new Date(str)
  return `${d.toLocaleDateString()} ${d.toLocaleTimeString()}`
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
