import { expect } from 'chai'
import selectByEnv from 'helpers/selectByEnv'

describe('selectByEnv', () => {
	it('returns undefined when no parameters are passed', () => {
		expect(selectByEnv()).to.be.undefined
	})

	it('returns undefined when mapping is incorrect', () => {
		expect(selectByEnv({ '': 'testing' })).to.be.undefined
	})

	it('returns the correct mapping when input is correct', () => {
		// NODE_ENV=hot is the default for all unit tests
		expect(selectByEnv({ 'hot': 'testing' })).to.equal('testing')
	})
})
