# Big List of Naughty Strings
The Big List of Naughty Strings is an evolving list of strings which have a high probability of causing issues when used as user-input data. This is intended for use in helping both automated and manual QA testing; useful for whenever your QA engineer [walks into a bar](http://www.sempf.net/post/On-Testing1.aspx).

## Why Test Naughty Strings?

Even multi-billion dollar companies with huge amounts of automated testing can't find every bad input. For example, look at what happens when you try to Tweet a [zero-width space](https://en.wikipedia.org/wiki/Zero-width_space) (U+200B) on Twitter:

![](http://i.imgur.com/HyDg2eV.gif)

Although this is not a malicious error, and typical users aren't Tweeting weird unicode, an "internal server error" for unexpected input is never a positive experience for the user, and may in fact be a symptom of deeper string-validation issues. The Big List of Naughty Strings is intended to help reveal such issues.

## Usage

`blns.txt` consists of newline-delimited strings and comments which are preceded with `#`. The comments divide the strings into sections for easy manual reading and copy/pasting into input forms. For those who want to access the strings programmatically, a `blns.json` file is provided containing an array with all the comments stripped out (the `scripts` folder contains a Python script used to generate the `blns.json`).

## Contributions

Feel free to send a pull request to add more strings, or additional sections. However, please do not send pull requests with very-long strings (255+ characters), as that makes the list much more difficult to view.

Likewise, please do not send pull requests which compromise *manual usability of the file*. This includes the [EICAR test string](https://en.wikipedia.org/wiki/EICAR_test_file), which can cause the file to be flagged by antivirus scanners, and files which alter the encoding of `blns.txt`. Also, do not send a null character (U+0000) string, as it [changes the file format on GitHub to binary](http://stackoverflow.com/a/19723302) and renders it unreadable in pull requests. Finally, when adding or removing a string please update all files when you perform a pull request.

## Disclaimer

The Big List of Naughty Strings is intended to be used *for software you own and manage*. Some of the Naughty Strings can indicate security vulnerabilities, and as a result using such strings with third-party software may be a crime. The maintainer is not responsible for any negative actions that result from the use of the list.

## Maintainer

* Max Woolf ([@minimaxir](https://twitter.com/minimaxir))
