# Java Verifier

Sometimes it's useful to assert that the code in two jars is identical. Why?
Maybe you're working on an open source Java project, and you realize that the code has never been formatted.
You run your favorite formatter and go to open a PR when you notice that you've changed thousands of lines.
Nobody wants to approve that PR because they can't be sure you aren't trying to sneak in something malicious.
However, with this tool you can assert the code in the jar before formatting is the same as the code in the jar after formatting.

### Installation

With a proper go installation you can simply `go install github.com/sosodev/java-verifier`. Otherwise you can download the binary for your system
from the releases tab.

### Usage

1) Disable compiler metadata by adding the `-g:none` flag when invoking `javac`. You can do this with gradle by adding `options.compilerArgs << '-g:none'` to the `JavaCompile` block in your `build.gradle` file.
2) Build the jar for the main branch and extract it somewhere (jars can be extracted as a normal zip file).
3) Build the jar for the comparison branch and extract it somewhere.
4) Execute the verifier like so `java-verifier <path-to-original-jar-directory> <path-to-comparison-jar-directory>`
