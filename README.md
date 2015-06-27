
# go-react

A series of **highly experimental** [GopherJS](gopherjs.org) bindings for 
[React.js](https://facebook.github.io/react/index.html).

Tests and benchmarks are absent until the proof of concept is more 
fleshed out, both in Go and React. The API is likely to change randomly 
and with great malice.

In short, until becnhmarks are implemented we can't know if this is 
feasible. So no, this is not ready for production. No, this is not ready 
for funzies. No, this is probably not ready for anything yet.

## Examples

A series of proof of concept examples, which match React docs 
examples/tutorials, can be found [in the \_examples 
directory](./_examples)

## Goals

1. Abstract away all `*js.Object` and `interface{}` tomfoolery from 
  GopherJS. Nice Go types and interfaces, all around.
2. Translate React-centric design elements to a more Go oriented design.
3. Provide a familiar API to create React components, while adhereing to 
  goal #2.

## Installation

Did you not read before? Don't use this!

## FAQ

### Q: Can i use this in the browser?

Yup! Just make sure a global react library object is available.

### Q: What about CommonJS?

Support for CommonJS _(Node/Webpack/etc)_ will be coming in the near 
future. Once the PoC is more fleshed out.

### Q: How's the performance?

A: Truthfully, no clue. I've not done any benchmarks at this point. I
expect some loss due to the overhead of GopherJS, but my hope is that if
we only create simple _(small surface)_ objects and pass them to React,
it won't know the difference.

If React doesn't experience anything different from a "normal" 
environment, we _may_ not have large losses.

Just make sure to follow GohperJS' [performance 
tips](https://github.com/gopherjs/gopherjs#performance-tips).

### Q: Didn't someone already try this?

A: Yup, and gave up. Unforutnately they removed their efforts from 
Github, so it's hard to say what exactly was done right, or wrong.

It also helps that i do not expect to succeed. I expect this to fail, 
somehow. And if this ends up working, win win right?

### Q: Why not a pure GopherJS Framework?

A: That would likely be far better, you are correct. Unfortunately 
GopherJS is quite young, and pure GopherJS frameworks are even more 
young. I wanted to use Go with a well vetted clientside framework, one 
that we trust and love. React certainly fits that description, though 
it's design may prove difficult to map GopherJS to.

Don't forget to checkout [Frameworks using 
GopherJS](https://github.com/gopherjs/gopherjs/wiki/Bindings#frameworks-using-gopherjs).

- https://github.com/gowade/wade
- https://github.com/go-humble/humble


### Q: Oh god why

A: I know right?

### Q: No seriously, why Go + React?

A: This project spawned as need to prove that it's not possible. I was
writing a pipeline for [FlowType](flowtype.org) and 
[BabelJS](babeljs.io), and adding in linting, formatting, and etc.  All 
with the goal of making writing JS nice and scalable for large projects 
_(mainly via FlowType's type checking)_.

After writing a rather complex and fragile pipeline to handle all of that 
stuff, it hit me. Why am i trying so hard to bend JavaScript to my needs?
FlowType is great and all, but it's a young project so there are a lot of 
rough edges with supporting tools.

Go on the otherhand, has all the nicities i love, it's a great language, 
and is what i'm already writing the backend in.

I'm not a believer in the holy grail of writing one language on all 
platforms. I have no motivations for that here. I simply want a nicely 
typed language with features and good tooling. Go achieves that.

### Q: But I love JS || Flow || Babel!

A: Great, stick with JS || Flow || Babel! You'll get better performance 
than you'll see here anyway :)
