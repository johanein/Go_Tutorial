module github.com/johanein/goTutorialHello

go 1.17

require (
	github.com/johanein/goTutorialGreetings v0.0.0-00010101000000-000000000000
	rsc.io/quote/v4 v4.0.1
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace github.com/johanein/goTutorialGreetings => ../greetings
