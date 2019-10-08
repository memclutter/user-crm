export function sortStrToArr(sort) {
  let srt = []

  if (sort.length !== 0) {
    const parts = sort.split(',')

    for (let i = 0; i < parts.length; i++) {
      let part = parts[i]
      if (part.startsWith('-')) {
        srt.push({name: part.substring(1, part.length), desc: true})
      } else {
        srt.push({name: part, desc: false})
      }
    }
  }
  return srt
}

export function sortArrToStr(value) {
  const sort = []
  for (let i = 0; i < value.length; i++) {
    let sign = value[i].desc ? '-' : ''
    sort.push(sign + value[i].name)
  }

  return sort.join(',')
}

