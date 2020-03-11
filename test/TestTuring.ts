const { test } = require('tape-modern');

test('Test turing engine outputs D(2,2)', t => {
	t.ok(true);
	t.ok(true, 'this time with an optional message');
	t.ok('not true, but truthy enough');

	t.equal(1 + 1, 2);
	t.equal(Math.max(1, 2, 3), 3);

	t.throws(() => {
		throw new Error('oh no!');
	}, /oh no!/);

	t.pass('this too shall pass');
});

test('these tests will not pass', t => {
	t.equal(42, '42');
	t.equal({}, {});
	t.fail('womp womp');
});

test.skip('this test will not run', t => {
	t.pass(false);
});