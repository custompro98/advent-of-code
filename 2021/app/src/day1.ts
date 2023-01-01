const counter = (input: number[]): number => {
  let numIncreases = 0

  // need at least two digits
  if ([0, 1].includes(input.length)) {
    return numIncreases
  }

  for (let i: number = 1; i < input.length; i++) {
    if (input[i] > input[i - 1]) {
      numIncreases++
    }
  }

  return numIncreases
}

const tripletCounter = (input: number[]): number => {
  let numIncreases = 0

  // need at least two triples
  if (input.length < 4) {
    return numIncreases
  }

  const triples = []

  for (let i: number = 2; i < input.length; i++) {
    triples.push(input[i] + input[i - 1] + input[i - 2])
  }

  for (let i: number = 1; i < triples.length; i++) {
    if (triples[i] > triples[i - 1]) {
      numIncreases++
    }
  }

  return numIncreases
}

export { counter, tripletCounter }
