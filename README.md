# Color Distance Visualizer
(in Go this time)

Saw the [Vidio library](https://github.com/AlexEidt/Vidio) on
Hacker News and liked its philosophy of using
FFmpeg subprocesses instead of working with `libav` directly
like [PyAV](https://pyav.org/docs/stable/). I had previously, naively thought
that "real" video processing couldn't be done over the command line and instead
necessitated working with a low-level library; I'm glad to have been proven wrong.
I've been meaning to try Go for a serious project for a while after getting halfway through
[a book on it](https://www.manning.com/books/go-in-practice),
and [my stupidity with writing efficient NumPy](github.com/supersonichub1/color-distance-visualizer-go)
made me feel as if I should "rewriting it in ~~Rust~~ Go" a try,
so I did.

## Usage
`git clone`, `cd`, `go get`, and `go run *.go`/`go build *.go`.

## Usage
```
Usage of color-distance-visualizer:
  -input string
        The video that we want to analyze.
  -output string
        Where we want to save the output.
  -show-unchanged-pixels
        For pixels that haven't changed, display the pixel from the orginal frame instead of black.
  -vcodec string
        What codec you want the output to be saved in. FFV1 is the recommended lossless codec. Using lossy codecs like MPEG-4 will result in significant loss of detail. (default: codec of input)
```

## Thoughts on Go
Go compiles really quickly and runs very fast, so there's basically no difference between
running a script and compiling code. Of course, I know that compilation
can take quite a while for large projects, so I guess I'll see how Go scales overtime.

I decided to go for an iterator approach using channels and a sweet little
struct. Bit disappointed that channels can't use tuples like return values can,
but I can get used to it. I think using a channel and goroutine is a bit overkill
just for some syntaxtic sweetness, but it doesn't seem to be harming performance.

Parts of Go's syntax are a bit awkward. In languages like C# and Java, you
write `int[] array`, but in Go, you write `array []int`. Since the latter
is muscle memory for me, I found myself getting syntax errors over and over again,
but I'm sure I'll get used to it just like switching between languages that do and
don't use semi-colons (by the way, I'm quite pleased that Go doesn't require semis!).

I feel like my re-writing of the visualizer is a lot easier to read, even though it's about
as long as the Python program in terms of lines. I think this is due to three things:
learning from my mistakes with writing the previous iteration,
Go forcing me to write simpler code compared to Python due to it's lack of fancy ways to manipulate
iterables, and Vidio having a significantly lighter API than that of PyAV plus NumPy.

If there's one thing that generally peeves me about Go, it's how awful `flag` is
compared to `argparse`. I can see why everyone and their mother brings in dependencies
when writing command-line tools due to how few creature comforts the standard library has. Why can't
I specify positional arguments using the libraries? It forces one to either make every input to 
their program a flag or forgo using functions like `Usage` and `PrintDefaults` and have to enforce types themselves.
I find it maddening that the Go devs implemented subcommands before they implemented positional arguments.
Why can't I state if a flag is optional required? Why can't I specify the name of my program so that user's don't
see `/tmp/go-build4204783715/b001/exe/main` when they run the program?

All in all, I like that churning through a 1080p 30-second video takes about
thirty seconds instead of two minutes. I plan to use Go and Vidio for any other
video processing projects I think of.
