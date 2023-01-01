enum Direction {
  None = "none",
  Up = "up",
  Down = "down",
  Forward = "forward",
}

const directionMap = new Map<string, Direction>([
  ["up", Direction.Up],
  ["down", Direction.Down],
  ["forward", Direction.Forward],
])

const parseDirection = (dirString: string): Direction => {
  if (!directionMap.has(dirString)) {
    return Direction.None
  }

  return directionMap.get(dirString) as Direction
}

type Command = {
  direction: Direction,
  distance: number,
}

const parseCommand = (command: string): Command => {
  const [directionString, distanceString] = command.split(" ")

  const direction = parseDirection(directionString)
  let distance = parseInt(distanceString)

  if (direction === Direction.Up) {
    distance *= -1
  }

  return {
    direction,
    distance
  }
}

type Result = {
  aim: number,
  x: number,
  y: number,
}

const newResult = (): Result => {
  return {
    aim: 0,
    x: 0,
    y: 0,
  }
}

const navigate = (input: string[]): number => {
  const result = input
    .map(parseCommand)
    .reduce((acc: Result, cmd) => {
      switch (cmd.direction) {
        case Direction.Up: {
          acc.y += cmd.distance
          break
        }
        case Direction.Down: {
          acc.y += cmd.distance
          break
        }
        case Direction.Forward: {
          acc.x += cmd.distance
          break
        }
      }

      return acc
    }, newResult())

  return result.x * result.y
}

const navigateAim = (input: string[]): number => {
  const result = input
    .map(parseCommand)
    .reduce((acc: Result, cmd) => {
      switch (cmd.direction) {
        case Direction.Up: {
          acc.aim += cmd.distance
          break
        }
        case Direction.Down: {
          acc.aim += cmd.distance
          break
        }
        case Direction.Forward: {
          acc.x += cmd.distance
          acc.y += acc.aim * cmd.distance
          break
        }
      }

      return acc
    }, newResult())

  return result.x * result.y
}

export { navigate, navigateAim }
