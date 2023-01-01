const world = 'world'

function hello(word: string = world): string {
    return `Hello ${word}!`
}

console.log(hello('docker'))
