# Big List of Naughty Strings
The Big List of Naughty Strings is an evolving list of strings which have a high probability of causing issues when used as user-input data. This is intended for use in helping both automated and manual QA testing; useful for whenever your QA engineer [walks into a bar](http://www.sempf.net/post/On-Testing1).

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

Additionally, the Big List of Naughty Strings is not a fully-comprehensive substitute for formal security/penetration testing for your service.

## Library / Packages

Various implementations of the Big List of Naughty Strings have made it to various package managers.  Those are maintained by outside parties, but can be found here:

| Library | Link |
| ------- | ---- |
| Node | https://www.npmjs.com/package/blns |
| Node | https://www.npmjs.com/package/big-list-of-naughty-strings |
| .NET | https://github.com/SimonCropp/NaughtyStrings |
| PHP | https://github.com/mattsparks/blns-php |
| C++  | https://github.com/eliabieri/blnscpp |

Please open a PR to list others.

## Maintainer/Creator

Max Woolf ([@minimaxir](https://twitter.com/minimaxir))

## Social Media Discussions

* June 10, 2015 [Hacker News]: [Show HN: Big List of Naughty Strings for testing user-input data](https://news.ycombinator.com/item?id=10035008)
* August 17, 2015 [Reddit]: [Big list of naughty strings.](https://www.reddit.com/r/programming/comments/3hdxqx/big_list_of_naughty_strings/)
* February 9, 2016 [Reddit]: [Big List of Naughty Strings](https://www.reddit.com/r/webdev/comments/44wc5b/big_list_of_naughty_strings/)
* January 15, 2017 [Hacker News]: [Naughty Strings: A list of strings likely to cause issues as user-input data](https://news.ycombinator.com/item?id=13406119)
* January 16, 2017 [Reddit]: [Naughty Strings: A list of strings likely to cause issues as user-input data](https://www.reddit.com/r/programming/comments/5o9inb/naughty_strings_a_list_of_strings_likely_to_cause/)
* November 16, 2018 [Hacker News]: [Big List of Naughty Strings](https://news.ycombinator.com/item?id=18466787)
* November 16, 2018 [Reddit]: [Naughty Strings - A list of strings which have a high probability of causing issues when used as user-input data](https://www.reddit.com/r/programming/comments/9xla2j/naughty_strings_a_list_of_strings_which_have_a/)

## License

MIT
