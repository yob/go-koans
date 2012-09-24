package go_koans

import "bytes"
import "io"

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		out.ReadFrom(in)

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		limited_in := io.LimitReader(in, 5)
		out.ReadFrom(limited_in)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
