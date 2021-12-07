import sum from '../sum'

describe('sum', () => {
    it('adds 1 + 2 together', () => {
        expect(sum(1, 2)).toBe(3)
    })
})
