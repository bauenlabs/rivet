// This package provides an efficient method by which strings can be
// concatenated. The code in this package is inspired directly by this article:
// http://herman.asia/efficient-string-concatenation-in-go.
package rivet

import "bytes"

// Takes an unlimited number of strings and concatenates them in the
// order that they were passed.
func Concat(strings ...string) string {
	var buffer bytes.Buffer
	for i := range strings {
		buffer.WriteString(strings[i])
	}

	return buffer.String()
}
