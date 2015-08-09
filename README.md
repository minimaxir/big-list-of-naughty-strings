# Big List of Naughty Strings
The Big List of Naughty Strings is a list of strings which have a high probability of causing issues when used as user-input data. This is intended for use in helping both automated and manual QA testing, for whenever your QA engineer [walks into a bar](http://www.sempf.net/post/On-Testing1.aspx).

The list consisted of newline-delimited strings and comments which preceded with `#`. The comments divide the strings into sections for easy manual reading.

## Why test naughty strings?

Even multi-billion dollar companies with infinite automated testing can't find every bad input. For example, look at what happens when you try to Tweet a zero-width space (U+200B):

![](http://i.imgur.com/AihSKMM.gif)

Although this is not a malicious error, an "internal server error" for unexpected input may be a symptom for deeper string-validation issues, wh