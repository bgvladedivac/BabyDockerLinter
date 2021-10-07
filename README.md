Design:
    There's a handler for each instruction. A handler is just a function with the concrete implementation of the instruction, for example EvaluateFromInstruction() is a handler, although in the name of the function handler is missing. A handler is more like an abstraction.

    For each handler we have a succes/non-success behavior test. The exception at this stage is the MAINTAINER handler, since if you have a MAITAINER you always get the same behavior ( the instruction is deprecated )

Linting rules:
    FROM must have a tag.
    EXPOSE port must be in supported port-range.
    MAINTAINER is flagged as deprecated.
    CMD must be matched only once in the file.

Run tests:
    cd /tests
    go test

Compile and run the binary:
    go build -o babylinter && ./babylinter [DIRECTORY_WHERE_DOCKERFILE_IS_LOCATED]

Contributions:
    Any contributions are welcomed.
