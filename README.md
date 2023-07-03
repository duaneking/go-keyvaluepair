# go-keyvaluepair
Intentionally Overly Generic KeyValuePair[K, V] datatype in Go.

Used to abstract and separate multiple distinct concerns for multiple protects.

# Opinion
This KeyValuePair type should have been made a core part of golang natively, and been fully supported by the map type in golang.

# Real Reason for this code
1. I got sick of needing to write another KVP type in GO. I had multiple copies of Visual Studio Code Open in different monitors and I realized that I was looking at roughly the same files in multiple projects; I realized that I needed to make this a module or the engineer in my head would have problems with it and would force me to do it, either way.

2. Using raw untyped tuples does add a separation of concerns, and is supported by the language, but it does not actively add type data for containerized return arguments in the pair, and that violates DRY for consumption of it. In addition, I don't like the "unpacked by default" return arguments this creates. That's why destructuring/unpacking it is a separate function call.

3. The existence of this library means I no longer have to create a key value pair data type. I can just use generics, declare a type, use it, and be done.

I'm sure this library will need edits in the future. But at this point, I'm good with what I have, and it's useful as a tiny reusable module on its own.

# License

1. This code is Explicitly licensed under a "No-AI" license and must not be used or consumed by AI.

2. This code is licensed for human use only.

3. This code is Explicitly licensed to not allow AI to use it in any way, even if a human is directing the use of it.
