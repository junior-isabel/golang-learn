module helloworld

go 1.17

require (
	github.com/junior-isabel/calculator v1.0.0
	github.com/junior-isabel/vector v1.0.0
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.3.7 // indirect
	rsc.io/sampler v1.99.99 // indirect
)

replace github.com/junior-isabel/calculator => ../calculator
replace github.com/junior-isabel/vector => ../vector